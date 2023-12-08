package test

import (
	"time"

	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
	v1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/pkg/test"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/agent"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/alerting"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/alarms/v1"
	endpointv1 "github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/endpoints/v1"
	"google.golang.org/protobuf/types/known/durationpb"
)

func init() {
	test.EnablePlugin(meta.ModeGateway, alerting.Scheme)

	alerting.DefaultDisconnectAlarm = func(clusterId string) *v1.AlertCondition {
		return &alertingv1.AlertCondition{
			Name:        "agent-disconnect",
			Description: "Alert when the downstream agent disconnects from the opni upstream",
			Labels:      []string{"agent-disconnect", "opni", "_default"},
			Severity:    alertingv1.OpniSeverity_Critical,
			AlertType: &alertingv1.AlertTypeDetails{
				Type: &alertingv1.AlertTypeDetails_System{
					System: &alertingv1.AlertConditionSystem{
						ClusterId: &corev1.Reference{Id: clusterId},
						Timeout:   durationpb.New(1 * time.Second),
					},
				},
			},
		}
	}

	alerting.DefaultCapabilityHealthAlarm = func(clusterId string) *v1.AlertCondition {
		return &alertingv1.AlertCondition{
			Name:        "agent-capability-unhealthy",
			Description: "Alert when some downstream agent capability becomes unhealthy",
			Labels:      []string{"agent-capability-health", "opni", "_default"},
			Severity:    alertingv1.OpniSeverity_Critical,
			AlertType: &alertingv1.AlertTypeDetails{
				Type: &alertingv1.AlertTypeDetails_DownstreamCapability{
					DownstreamCapability: &alertingv1.AlertConditionDownstreamCapability{
						ClusterId:       &corev1.Reference{Id: clusterId},
						CapabilityState: alerting.ListBadDefaultStatuses(),
						For:             durationpb.New(1 * time.Second),
					},
				},
			},
		}
	}

	alerting.SyncInterval = time.Second * 1
	alerting.ForceSyncInterval = time.Second * 60

	alarms.DisconnectStreamEvaluateInterval = time.Second * 1
	alarms.CapabilityStreamEvaluateInterval = time.Minute * 100
	alarms.CortexStreamEvaluateInterval = time.Second * 1
	test.EnablePlugin(meta.ModeAgent, agent.Scheme)
	endpointv1.RetryTestEdnpoint = time.Millisecond * 50

	agent.RuleSyncInterval = time.Second * 1
}
