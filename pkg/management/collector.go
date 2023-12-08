package management

import (
	"context"
	"strings"
	"time"

	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	managementv1 "github.com/open-panoptes/opni/pkg/apis/management/v1"
	"github.com/open-panoptes/opni/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/samber/lo"
)

var (
	clusterInfo = prometheus.NewDesc(
		"opni_cluster_info",
		"Cluster information",
		[]string{metrics.LabelImpersonateAs, "friendly_name"},
		prometheus.Labels{},
	)
	agentUp = prometheus.NewDesc(
		"opni_agent_up",
		"Agent connection status",
		[]string{metrics.LabelImpersonateAs},
		prometheus.Labels{},
	)
	agentReady = prometheus.NewDesc(
		"opni_agent_ready",
		"Agent readiness status",
		[]string{metrics.LabelImpersonateAs, "conditions"},
		prometheus.Labels{},
	)
	agentSummary = prometheus.NewDesc(
		"opni_agent_status_summary",
		"Agent status summary",
		[]string{metrics.LabelImpersonateAs, "summary"},
		prometheus.Labels{},
	)
)

func (s *Server) Describe(c chan<- *prometheus.Desc) {
	c <- clusterInfo
	c <- agentUp
	c <- agentReady
	c <- agentSummary
}

func (s *Server) Collect(c chan<- prometheus.Metric) {
	ctx, ca := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer ca()
	clusters, err := s.ListClusters(ctx, &managementv1.ListClustersRequest{})
	if err != nil {
		return
	}
	for _, cluster := range clusters.Items {
		var friendlyName = cluster.Id
		labels := cluster.GetLabels()
		if nameLabel, ok := labels[corev1.NameLabel]; ok {
			friendlyName = nameLabel
		} else if legacyNameLabel, ok := labels[corev1.LegacyNameLabel]; ok {
			friendlyName = legacyNameLabel
		}
		c <- prometheus.MustNewConstMetric(
			clusterInfo,
			prometheus.GaugeValue,
			1,
			cluster.Id,
			friendlyName,
		)

		var connected, ready float64
		var conditions, summary string
		if hs, err := s.GetClusterHealthStatus(ctx, &corev1.Reference{Id: cluster.Id}); err == nil {
			connected = lo.Ternary[float64](hs.Status.Connected, 1, 0)
			// todo: possible bug here? hs.Health==nil
			ready = lo.Ternary[float64](hs.GetHealth().GetReady(), 1, 0)
			conditions = strings.Join(hs.GetHealth().GetConditions(), ",")
			summary = hs.Summary()
		}
		c <- prometheus.MustNewConstMetric(
			agentUp,
			prometheus.GaugeValue,
			connected,
			cluster.Id,
		)
		c <- prometheus.MustNewConstMetric(
			agentReady,
			prometheus.GaugeValue,
			ready,
			cluster.Id,
			conditions,
		)
		c <- prometheus.MustNewConstMetric(
			agentSummary,
			prometheus.GaugeValue,
			1,
			cluster.Id,
			summary,
		)
	}
}
