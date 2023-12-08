package templates_test

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-panoptes/opni/pkg/alerting/drivers/config"
	"github.com/open-panoptes/opni/pkg/alerting/interfaces"
	"github.com/open-panoptes/opni/pkg/alerting/templates"
	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	"github.com/open-panoptes/opni/pkg/util"

	"google.golang.org/protobuf/types/known/durationpb"

	"text/template"

	"github.com/open-panoptes/opni/pkg/alerting/message"
	amtemplate "github.com/prometheus/alertmanager/template"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/promql"
	promtemplate "github.com/prometheus/prometheus/template"
)

func init() {
	templates.RegisterNewAlertManagerDefaults(amtemplate.DefaultFuncs, templates.DefaultTemplateFuncs)
}

type TestMessage interface {
	interfaces.Routable
	Sanitize()
	Validate() error
}

var _ = Describe("Prometheus fingerprint templating", Label("unit"), func() {
	When("we use a promteheus template expander we should get back a fingerprint", func() {
		It("should create a valid template expander mock", func() {
			prometheusAlert := &alertingv1.AlertCondition{
				Name:        "test header 1",
				Description: "test header 2",
				Id:          uuid.New().String(),
				Severity:    alertingv1.OpniSeverity_Error,
				AlertType: &alertingv1.AlertTypeDetails{
					Type: &alertingv1.AlertTypeDetails_PrometheusQuery{
						PrometheusQuery: &alertingv1.AlertConditionPrometheusQuery{
							ClusterId: &corev1.Reference{Id: uuid.New().String()},
							Query:     "up > 0",
						},
					},
				},
			}
			prometheusFingerprintTemplate := prometheusAlert.GetRoutingAnnotations()[message.NotificationPropertyFingerprint]
			fingerprintTs := float64(time.Now().UnixMilli()) / 1000
			scenarios := []scenario{
				{
					text:   "{{ \"ALERTS_FOR_STATE{opni_uuid=\\\"%s\\\"} OR on() vector(0))\" | query | first | value | printf \"%.0f\" }}",
					output: "0",
					queryResult: promql.Vector{
						{
							T: time.Now().Unix(),
							F: 0,
						},
					},
				},
				{
					text:   prometheusFingerprintTemplate,
					output: fmt.Sprintf("%.0f", fingerprintTs),
					queryResult: promql.Vector{
						{
							Metric: labels.Labels{
								{
									Name:  "__name__",
									Value: "ALERTS_FOR_STATE",
								},
							},
							T: time.Now().Unix(),
							F: fingerprintTs,
						},
					},
				},
			}
			for _, s := range scenarios {
				queryFunc := func(_ context.Context, _ string, _ time.Time) (promql.Vector, error) {
					return s.queryResult, nil
				}
				expander := promtemplate.NewTemplateExpander(
					context.Background(),
					s.text,
					"test",
					s.input,
					0,
					queryFunc,
					util.Must(url.Parse("http://localhost:9093")),
					s.options,
				)
				result, err := expander.Expand()
				Expect(err).ToNot(HaveOccurred())
				Expect(result).To(Equal(s.output))
			}
		})
	})
})

var _ = DescribeTable("Message templating",
	func(incomingMsg TestMessage, status string, headerContains, bodyContains []string) {
		Expect(incomingMsg).ToNot(BeNil())
		incomingMsg.Sanitize()
		Expect(incomingMsg.Validate()).ToNot(HaveOccurred())

		msg := &config.WebhookMessage{
			Receiver: "test",
			Status:   "firing",
			Alerts: config.Alerts{
				{
					Status:      status,
					Labels:      incomingMsg.GetRoutingLabels(),
					Annotations: incomingMsg.GetRoutingAnnotations(),
					StartsAt:    time.Now(),
				},
			},
			Version:         "v4",
			ExternalURL:     "http://localhost:9093",
			TruncatedAlerts: 0,
		}

		headerTmpl, err := template.New("").Funcs(
			template.FuncMap(amtemplate.DefaultFuncs),
		).Parse(templates.HeaderTemplate())
		Expect(err).ToNot(HaveOccurred())

		bodyTmpl, err := template.New("").Funcs(
			template.FuncMap(amtemplate.DefaultFuncs),
		).Parse(templates.BodyTemplate())
		Expect(err).ToNot(HaveOccurred())

		var (
			b1 bytes.Buffer
			b2 bytes.Buffer
		)
		err = headerTmpl.Execute(&b1, msg)
		Expect(err).ToNot(HaveOccurred())

		err = bodyTmpl.Execute(&b2, msg)
		Expect(err).ToNot(HaveOccurred())

		s1, s2 := b1.String(), b2.String()

		for _, s := range headerContains {
			Expect(s1).To(ContainSubstring(s))
		}

		for _, s := range bodyContains {
			Expect(s2).To(ContainSubstring(s))
		}

		By("verifying the cluster name is included in the message if it exists")
		if clusterName, ok := incomingMsg.GetRoutingAnnotations()[message.NotificationContentClusterName]; ok {
			Expect(b1.String()).To(ContainSubstring(clusterName))
		} else if clusterId, ok := incomingMsg.GetRoutingLabels()[message.NotificationPropertyClusterId]; ok {
			Expect(b1.String()).To(ContainSubstring(clusterId))
		}
		By("verifying the alarm name is included in the message if it exists")
		if alarmName, ok := incomingMsg.GetRoutingAnnotations()[message.NotificationContentAlarmName]; ok {
			Expect(b1.String()).To(ContainSubstring(alarmName))
		}

		By("verifying that we output pretty timestamps in messages")
		tsFunc, ok := amtemplate.DefaultFuncs["formatTime"]
		Expect(ok).To(BeTrue())

		v := reflect.ValueOf(tsFunc)
		Expect(v.Kind()).To(Equal(reflect.Func))

		prettyTs := v.Call([]reflect.Value{reflect.ValueOf(msg.Alerts[0].StartsAt)})
		Expect(prettyTs).To(HaveLen(2))

		Expect(prettyTs[1].IsNil()).To(BeTrue())
		Expect(prettyTs[0].Kind()).To(Equal(reflect.String))
		Expect(prettyTs[0].String()).ToNot(HaveLen(0))
		Expect(strings.Contains(s1, prettyTs[0].String()) || strings.Contains(s2, prettyTs[0].String())).To(BeTrue())
	},
	Entry(
		"firing alarm uses user's custom title and body",
		&alertingv1.AlertCondition{
			Name:        "condition 1",
			Description: "condition 1 description",
			Severity:    alertingv1.OpniSeverity_Info,
			AlertType: &alertingv1.AlertTypeDetails{
				Type: &alertingv1.AlertTypeDetails_System{
					System: &alertingv1.AlertConditionSystem{
						ClusterId: &corev1.Reference{Id: uuid.New().String()},
						Timeout:   durationpb.New(10 * time.Minute),
					},
				},
			},
			AttachedEndpoints: &alertingv1.AttachedEndpoints{
				Items: []*alertingv1.AttachedEndpoint{
					{EndpointId: uuid.New().String()},
				},
				Details: &alertingv1.EndpointImplementation{
					Title: "hello",
					Body:  "world",
				},
			},
			Annotations: map[string]string{
				message.NotificationContentClusterName: "some-ridiculous-cluster-name",
			},
		},
		"firing",
		[]string{alertingv1.OpniSeverity_Info.String(), "Alarm", "condition 1"},
		[]string{"hello", "world"},
	),
	Entry(
		"firing alarm falls back to name & description",
		&alertingv1.AlertCondition{
			Name:        "condition 1",
			Description: "condition 1 description",
			Severity:    alertingv1.OpniSeverity_Warning,
			AlertType: &alertingv1.AlertTypeDetails{
				Type: &alertingv1.AlertTypeDetails_System{
					System: &alertingv1.AlertConditionSystem{
						ClusterId: &corev1.Reference{Id: uuid.New().String()},
						Timeout:   durationpb.New(10 * time.Minute),
					},
				},
			},
		},
		"firing",
		[]string{"condition 1", "Warning", alertingv1.OpniSeverity_Warning.String(), "firing"},
		[]string{"condition 1", "condition 1 description"},
	),
	Entry(
		"opni alarm is resolved",
		&alertingv1.AlertCondition{
			Name:        "test header 2",
			Description: "body 2",
			Severity:    alertingv1.OpniSeverity_Error,
			AlertType: &alertingv1.AlertTypeDetails{
				Type: &alertingv1.AlertTypeDetails_System{
					System: &alertingv1.AlertConditionSystem{
						ClusterId: &corev1.Reference{Id: uuid.New().String()},
						Timeout:   durationpb.New(10 * time.Minute),
					},
				},
			},
		},
		"resolved",
		[]string{"test header 2", alertingv1.OpniSeverity_Error.String(), "Alarm", "resolved"},
		[]string{"test header 2", "body 2"},
	),
	Entry(
		"critical notification",
		&alertingv1.Notification{
			Title: "notification title",
			Body:  "notification body",
			Properties: map[string]string{
				message.NotificationPropertySeverity: "Critical",
			},
		},
		"firing",
		[]string{"Critical", "Notification", "notification title"},
		[]string{"notification body"},
	),
)

// From : github.com/prometheus/prometheus/template/template_test.go

type scenario struct {
	text        string
	output      string
	input       interface{}
	options     []string
	queryResult promql.Vector
}
