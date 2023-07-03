// Code generated by internal/codegen. DO NOT EDIT.

// Code generated by cli_gen.go DO NOT EDIT.
// source: github.com/rancher/opni/internal/cortex/config/validation/limits.proto

package validation

import (
	flagutil "github.com/rancher/opni/pkg/util/flagutil"
	pflag "github.com/spf13/pflag"
	proto "google.golang.org/protobuf/proto"
	strings "strings"
	time "time"
)

func (in *Limits) DeepCopyInto(out *Limits) {
	proto.Merge(in, out)
}

func (in *Limits) DeepCopy() *Limits {
	return proto.Clone(in).(*Limits)
}

func (in *RelabelConfig) DeepCopyInto(out *RelabelConfig) {
	proto.Merge(in, out)
}

func (in *RelabelConfig) DeepCopy() *RelabelConfig {
	return proto.Clone(in).(*RelabelConfig)
}

func (in *Limits) FlagSet(prefix ...string) *pflag.FlagSet {
	fs := pflag.NewFlagSet("Limits", pflag.ExitOnError)
	fs.SortFlags = true
	fs.Float64Var(&in.IngestionRate, strings.Join(append(prefix, "ingestion-rate"), "."), 25000, "Per-user ingestion rate limit in samples per second.")
	fs.StringVar(&in.IngestionRateStrategy, strings.Join(append(prefix, "ingestion-rate-strategy"), "."), "local", "Whether the ingestion rate limit should be applied individually to each distributor instance (local), or evenly shared across the cluster (global).")
	fs.Int32Var(&in.IngestionBurstSize, strings.Join(append(prefix, "ingestion-burst-size"), "."), 50000, "Per-user allowed ingestion burst size (in number of samples).")
	fs.BoolVar(&in.AcceptHaSamples, strings.Join(append(prefix, "accept-ha-samples"), "."), false, "Flag to enable, for all users, handling of samples with external labels identifying replicas in an HA Prometheus setup.")
	fs.StringVar(&in.HaClusterLabel, strings.Join(append(prefix, "ha-cluster-label"), "."), "cluster", "Prometheus label to look for in samples to identify a Prometheus HA cluster.")
	fs.StringVar(&in.HaReplicaLabel, strings.Join(append(prefix, "ha-replica-label"), "."), "__replica__", "Prometheus label to look for in samples to identify a Prometheus HA replica.")
	fs.Int32Var(&in.HaMaxClusters, strings.Join(append(prefix, "ha-max-clusters"), "."), 0, "Maximum number of clusters that HA tracker will keep track of for single user. 0 to disable the limit.")
	fs.StringSliceVar(&in.DropLabels, strings.Join(append(prefix, "drop-labels"), "."), []string{}, "This flag can be used to specify label names that to drop during sample ingestion within the distributor and can be repeated in order to drop multiple labels.")
	fs.Int32Var(&in.MaxLabelNameLength, strings.Join(append(prefix, "max-label-name-length"), "."), 1024, "Maximum length accepted for label names")
	fs.Int32Var(&in.MaxLabelValueLength, strings.Join(append(prefix, "max-label-value-length"), "."), 2048, "Maximum length accepted for label value. This setting also applies to the metric name")
	fs.Int32Var(&in.MaxLabelNamesPerSeries, strings.Join(append(prefix, "max-label-names-per-series"), "."), 30, "Maximum number of label names per series.")
	fs.Int32Var(&in.MaxLabelsSizeBytes, strings.Join(append(prefix, "max-labels-size-bytes"), "."), 0, "Maximum combined size in bytes of all labels and label values accepted for a series. 0 to disable the limit.")
	fs.Int32Var(&in.MaxMetadataLength, strings.Join(append(prefix, "max-metadata-length"), "."), 1024, "Maximum length accepted for metric metadata. Metadata refers to Metric Name, HELP and UNIT.")
	fs.BoolVar(&in.RejectOldSamples, strings.Join(append(prefix, "reject-old-samples"), "."), false, "Reject old samples.")
	fs.Var(flagutil.DurationpbValue(336*time.Hour, &in.RejectOldSamplesMaxAge), strings.Join(append(prefix, "reject-old-samples-max-age"), "."), "Maximum accepted sample age before rejecting.")
	fs.Var(flagutil.DurationpbValue(10*time.Minute, &in.CreationGracePeriod), strings.Join(append(prefix, "creation-grace-period"), "."), "Duration which table will be created/deleted before/after it's needed; we won't accept sample from before this time.")
	fs.BoolVar(&in.EnforceMetadataMetricName, strings.Join(append(prefix, "enforce-metadata-metric-name"), "."), true, "Enforce every metadata has a metric name.")
	fs.BoolVar(&in.EnforceMetricName, strings.Join(append(prefix, "enforce-metric-name"), "."), true, "Enforce every sample has a metric name.")
	fs.Int32Var(&in.IngestionTenantShardSize, strings.Join(append(prefix, "ingestion-tenant-shard-size"), "."), 0, "The default tenant's shard size when the shuffle-sharding strategy is used. Must be set both on ingesters and distributors. When this setting is specified in the per-tenant overrides, a value of 0 disables shuffle sharding for the tenant.")
	fs.Int32Var(&in.MaxExemplars, strings.Join(append(prefix, "max-exemplars"), "."), 0, "Enables support for exemplars in TSDB and sets the maximum number that will be stored. less than zero means disabled. If the value is set to zero, cortex will fallback to blocks-storage.tsdb.max-exemplars value.")
	fs.Int32Var(&in.MaxSeriesPerQuery, strings.Join(append(prefix, "max-series-per-query"), "."), 100000, "The maximum number of series for which a query can fetch samples from each ingester. This limit is enforced only in the ingesters (when querying samples not flushed to the storage yet) and it's a per-instance limit. This limit is ignored when running the Cortex blocks storage. When running Cortex with blocks storage use -querier.max-fetched-series-per-query limit instead.")
	fs.Int32Var(&in.MaxSeriesPerUser, strings.Join(append(prefix, "max-series-per-user"), "."), 5000000, "The maximum number of active series per user, per ingester. 0 to disable.")
	fs.Int32Var(&in.MaxSeriesPerMetric, strings.Join(append(prefix, "max-series-per-metric"), "."), 50000, "The maximum number of active series per metric name, per ingester. 0 to disable.")
	fs.Int32Var(&in.MaxGlobalSeriesPerUser, strings.Join(append(prefix, "max-global-series-per-user"), "."), 0, "The maximum number of active series per user, across the cluster before replication. 0 to disable. Supported only if -distributor.shard-by-all-labels is true.")
	fs.Int32Var(&in.MaxGlobalSeriesPerMetric, strings.Join(append(prefix, "max-global-series-per-metric"), "."), 0, "The maximum number of active series per metric name, across the cluster before replication. 0 to disable.")
	fs.Int32Var(&in.MaxMetadataPerUser, strings.Join(append(prefix, "max-metadata-per-user"), "."), 8000, "The maximum number of active metrics with metadata per user, per ingester. 0 to disable.")
	fs.Int32Var(&in.MaxMetadataPerMetric, strings.Join(append(prefix, "max-metadata-per-metric"), "."), 10, "The maximum number of metadata per metric, per ingester. 0 to disable.")
	fs.Int32Var(&in.MaxGlobalMetadataPerUser, strings.Join(append(prefix, "max-global-metadata-per-user"), "."), 0, "The maximum number of active metrics with metadata per user, across the cluster. 0 to disable. Supported only if -distributor.shard-by-all-labels is true.")
	fs.Int32Var(&in.MaxGlobalMetadataPerMetric, strings.Join(append(prefix, "max-global-metadata-per-metric"), "."), 0, "The maximum number of metadata per metric, across the cluster. 0 to disable.")
	fs.Var(flagutil.DurationpbValue(0, &in.OutOfOrderTimeWindow), strings.Join(append(prefix, "out-of-order-time-window"), "."), "[Experimental] Configures the allowed time window for ingestion of out-of-order samples. Disabled (0s) by default.")
	fs.Int32Var(&in.MaxFetchedChunksPerQuery, strings.Join(append(prefix, "max-fetched-chunks-per-query"), "."), 2000000, "Maximum number of chunks that can be fetched in a single query from ingesters and long-term storage. This limit is enforced in the querier, ruler and store-gateway. 0 to disable.")
	fs.Int32Var(&in.MaxFetchedSeriesPerQuery, strings.Join(append(prefix, "max-fetched-series-per-query"), "."), 0, "The maximum number of unique series for which a query can fetch samples from each ingesters and blocks storage. This limit is enforced in the querier, ruler and store-gateway. 0 to disable")
	fs.Int32Var(&in.MaxFetchedChunkBytesPerQuery, strings.Join(append(prefix, "max-fetched-chunk-bytes-per-query"), "."), 0, "Deprecated (use max-fetched-data-bytes-per-query instead): The maximum size of all chunks in bytes that a query can fetch from each ingester and storage. This limit is enforced in the querier, ruler and store-gateway. 0 to disable.")
	fs.Int32Var(&in.MaxFetchedDataBytesPerQuery, strings.Join(append(prefix, "max-fetched-data-bytes-per-query"), "."), 0, "The maximum combined size of all data that a query can fetch from each ingester and storage. This limit is enforced in the querier and ruler for `query`, `query_range` and `series` APIs. 0 to disable.")
	fs.Var(flagutil.DurationpbValue(0, &in.MaxQueryLookback), strings.Join(append(prefix, "max-query-lookback"), "."), "Limit how long back data (series and metadata) can be queried, up until <lookback> duration ago. This limit is enforced in the query-frontend, querier and ruler. If the requested time range is outside the allowed range, the request will not fail but will be manipulated to only query data within the allowed time range. 0 to disable.")
	fs.Var(flagutil.DurationpbValue(0, &in.MaxQueryLength), strings.Join(append(prefix, "max-query-length"), "."), "Limit the query time range (end - start time). This limit is enforced in the query-frontend (on the received query) and in the querier (on the query possibly split by the query-frontend). 0 to disable.")
	fs.Int32Var(&in.MaxQueryParallelism, strings.Join(append(prefix, "max-query-parallelism"), "."), 14, "Maximum number of split queries will be scheduled in parallel by the frontend.")
	fs.Var(flagutil.DurationpbValue(1*time.Minute, &in.MaxCacheFreshness), strings.Join(append(prefix, "max-cache-freshness"), "."), "Most recent allowed cacheable result per-tenant, to prevent caching very recent results that might still be in flux.")
	fs.Float64Var(&in.MaxQueriersPerTenant, strings.Join(append(prefix, "max-queriers-per-tenant"), "."), 0, "Maximum number of queriers that can handle requests for a single tenant. If set to 0 or value higher than number of available queriers, *all* queriers will handle requests for the tenant. If the value is < 1, it will be treated as a percentage and the gets a percentage of the total queriers. Each frontend (or query-scheduler, if used) will select the same set of queriers for the same tenant (given that all queriers are connected to all frontends / query-schedulers). This option only works with queriers connecting to the query-frontend / query-scheduler, not when using downstream URL.")
	fs.Int32Var(&in.MaxOutstandingRequestsPerTenant, strings.Join(append(prefix, "max-outstanding-requests-per-tenant"), "."), 100, "Maximum number of outstanding requests per tenant per request queue (either query frontend or query scheduler); requests beyond this error with HTTP 429.")
	fs.Var(flagutil.DurationpbValue(0, &in.RulerEvaluationDelayDuration), strings.Join(append(prefix, "ruler-evaluation-delay-duration"), "."), "Duration to delay the evaluation of rules to ensure the underlying metrics have been pushed to Cortex.")
	fs.Int32Var(&in.RulerTenantShardSize, strings.Join(append(prefix, "ruler-tenant-shard-size"), "."), 0, "The default tenant's shard size when the shuffle-sharding strategy is used by ruler. When this setting is specified in the per-tenant overrides, a value of 0 disables shuffle sharding for the tenant.")
	fs.Int32Var(&in.RulerMaxRulesPerRuleGroup, strings.Join(append(prefix, "ruler-max-rules-per-rule-group"), "."), 0, "Maximum number of rules per rule group per-tenant. 0 to disable.")
	fs.Int32Var(&in.RulerMaxRuleGroupsPerTenant, strings.Join(append(prefix, "ruler-max-rule-groups-per-tenant"), "."), 0, "Maximum number of rule groups per-tenant. 0 to disable.")
	fs.Float64Var(&in.StoreGatewayTenantShardSize, strings.Join(append(prefix, "store-gateway-tenant-shard-size"), "."), 0, "The default tenant's shard size when the shuffle-sharding strategy is used. Must be set when the store-gateway sharding is enabled with the shuffle-sharding strategy. When this setting is specified in the per-tenant overrides, a value of 0 disables shuffle sharding for the tenant. If the value is < 1 the shard size will be a percentage of the total store-gateways.")
	fs.Int32Var(&in.MaxDownloadedBytesPerRequest, strings.Join(append(prefix, "max-downloaded-bytes-per-request"), "."), 0, "The maximum number of data bytes to download per gRPC request in Store Gateway, including Series/LabelNames/LabelValues requests. 0 to disable.")
	fs.Var(flagutil.DurationpbValue(0, &in.CompactorBlocksRetentionPeriod), strings.Join(append(prefix, "compactor-blocks-retention-period"), "."), "Delete blocks containing samples older than the specified retention period. 0 to disable.")
	fs.Int32Var(&in.CompactorTenantShardSize, strings.Join(append(prefix, "compactor-tenant-shard-size"), "."), 0, "The default tenant's shard size when the shuffle-sharding strategy is used by the compactor. When this setting is specified in the per-tenant overrides, a value of 0 disables shuffle sharding for the tenant.")
	fs.Var(flagutil.IPNetSliceValue(nil, &in.AlertmanagerReceiversFirewallBlockCidrNetworks), strings.Join(append(prefix, "alertmanager-receivers-firewall-block-cidr-networks"), "."), "Comma-separated list of network CIDRs to block in Alertmanager receiver integrations.")
	fs.BoolVar(&in.AlertmanagerReceiversFirewallBlockPrivateAddresses, strings.Join(append(prefix, "alertmanager-receivers-firewall-block-private-addresses"), "."), false, "True to block private and local addresses in Alertmanager receiver integrations. It blocks private addresses defined by  RFC 1918 (IPv4 addresses) and RFC 4193 (IPv6 addresses), as well as loopback, local unicast and local multicast addresses.")
	fs.Float64Var(&in.AlertmanagerNotificationRateLimit, strings.Join(append(prefix, "alertmanager-notification-rate-limit"), "."), 0, "Per-user rate limit for sending notifications from Alertmanager in notifications/sec. 0 = rate limit disabled. Negative value = no notifications are allowed.")
	fs.Var(flagutil.StringToFloat64Value(map[string]float64{}, &in.AlertmanagerNotificationRateLimitPerIntegration), strings.Join(append(prefix, "alertmanager-notification-rate-limit-per-integration"), "."), "Per-integration notification rate limits. Value is a map, where each key is integration name and value is a rate-limit (float). On command line, this map is given in JSON format. Rate limit has the same meaning as -alertmanager.notification-rate-limit, but only applies for specific integration. Allowed integration names: webhook, email, pagerduty, opsgenie, wechat, slack, victorops, pushover, sns.")
	fs.Int32Var(&in.AlertmanagerMaxConfigSizeBytes, strings.Join(append(prefix, "alertmanager-max-config-size-bytes"), "."), 0, "Maximum size of configuration file for Alertmanager that tenant can upload via Alertmanager API. 0 = no limit.")
	fs.Int32Var(&in.AlertmanagerMaxTemplatesCount, strings.Join(append(prefix, "alertmanager-max-templates-count"), "."), 0, "Maximum number of templates in tenant's Alertmanager configuration uploaded via Alertmanager API. 0 = no limit.")
	fs.Int32Var(&in.AlertmanagerMaxTemplateSizeBytes, strings.Join(append(prefix, "alertmanager-max-template-size-bytes"), "."), 0, "Maximum size of single template in tenant's Alertmanager configuration uploaded via Alertmanager API. 0 = no limit.")
	fs.Int32Var(&in.AlertmanagerMaxDispatcherAggregationGroups, strings.Join(append(prefix, "alertmanager-max-dispatcher-aggregation-groups"), "."), 0, "Maximum number of aggregation groups in Alertmanager's dispatcher that a tenant can have. Each active aggregation group uses single goroutine. When the limit is reached, dispatcher will not dispatch alerts that belong to additional aggregation groups, but existing groups will keep working properly. 0 = no limit.")
	fs.Int32Var(&in.AlertmanagerMaxAlertsCount, strings.Join(append(prefix, "alertmanager-max-alerts-count"), "."), 0, "Maximum number of alerts that a single user can have. Inserting more alerts will fail with a log message and metric increment. 0 = no limit.")
	fs.Int32Var(&in.AlertmanagerMaxAlertsSizeBytes, strings.Join(append(prefix, "alertmanager-max-alerts-size-bytes"), "."), 0, "Maximum total size of alerts that a single user can have, alert size is the sum of the bytes of its labels, annotations and generatorURL. Inserting more alerts will fail with a log message and metric increment. 0 = no limit.")
	return fs
}