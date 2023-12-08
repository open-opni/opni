package management_test

import (
	"context"
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	"github.com/open-panoptes/opni/pkg/metrics"
	"github.com/open-panoptes/opni/pkg/plugins"
	"github.com/prometheus/client_golang/prometheus"
)

func descriptorString(fqName, help string, constLabels []string, varLabels []prometheus.ConstrainedLabel) string {
	var vlStrings []string
	for _, varLabel := range varLabels {
		vlStrings = append(vlStrings, varLabel.Name)
	}

	return fmt.Sprintf(
		"Desc{fqName: %q, help: %q, constLabels: {%s}, variableLabels: {%s}}",
		fqName,
		help,
		strings.Join(constLabels, ","),
		strings.Join(vlStrings, ","),
	)
}

var _ = Describe("Collector", Ordered, Label("unit"), func() {
	var tv *testVars
	BeforeAll(setupManagementServer(&tv, plugins.NoopLoader))

	When("no clusters are present", func() {
		It("should collect descriptors but no metrics", func() {
			descs := make(chan *prometheus.Desc, 100)
			tv.ifaces.collector.Describe(descs)
			Eventually(descs).Should(HaveLen(4))
			Consistently(descs).Should(HaveLen(4))
			metrics := make(chan prometheus.Metric, 100)
			tv.ifaces.collector.Collect(metrics)
			Consistently(metrics).Should(BeEmpty())
		})
	})
	When("clusters are present", func() {
		It("should collect metrics for each cluster", func() {
			tv.storageBackend.CreateCluster(context.Background(), &corev1.Cluster{
				Id: "cluster-1",
				Metadata: &corev1.ClusterMetadata{
					Labels:       map[string]string{corev1.NameLabel: "cluster-1"},
					Capabilities: []*corev1.ClusterCapability{{Name: "test"}},
				},
			})
			tv.storageBackend.CreateCluster(context.Background(), &corev1.Cluster{
				Id: "cluster-2",
				Metadata: &corev1.ClusterMetadata{
					Labels:       map[string]string{corev1.NameLabel: "cluster-2"},
					Capabilities: []*corev1.ClusterCapability{{Name: "test"}},
				},
			})

			c := make(chan *prometheus.Desc, 100)
			tv.ifaces.collector.Describe(c)
			Expect(c).To(HaveLen(4))
			descs := make([]string, 0, 4)
			for i := 0; i < 4; i++ {
				descs = append(descs, (<-c).String())
			}
			Expect(descs).To(ConsistOf(
				descriptorString(
					"opni_cluster_info",
					"Cluster information",
					[]string{},
					[]prometheus.ConstrainedLabel{
						{
							Name: metrics.LabelImpersonateAs,
						},
						{
							Name: "friendly_name",
						},
					},
				),
				descriptorString(
					"opni_agent_up",
					"Agent connection status",
					[]string{},
					[]prometheus.ConstrainedLabel{
						{
							Name: metrics.LabelImpersonateAs,
						},
					},
				),
				descriptorString(
					"opni_agent_ready",
					"Agent readiness status",
					[]string{},
					[]prometheus.ConstrainedLabel{
						{
							Name: metrics.LabelImpersonateAs,
						},
						{
							Name: "conditions",
						},
					},
				),
				descriptorString(
					"opni_agent_status_summary",
					"Agent status summary",
					[]string{},
					[]prometheus.ConstrainedLabel{
						{
							Name: metrics.LabelImpersonateAs,
						},
						{
							Name: "summary",
						},
					},
				),
			))

			metrics := make(chan prometheus.Metric, 100)
			tv.ifaces.collector.Collect(metrics)
			Expect(metrics).To(Receive())
			Expect(metrics).To(Receive())
		})
	})
})
