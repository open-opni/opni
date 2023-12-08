package alerting

import (
	"context"
	"crypto/tls"

	"github.com/open-panoptes/opni/pkg/agent"
	"github.com/open-panoptes/opni/pkg/config/v1beta1"
	"github.com/open-panoptes/opni/pkg/management"
	"github.com/open-panoptes/opni/pkg/metrics/collector"
	"github.com/open-panoptes/opni/pkg/storage"
	"github.com/open-panoptes/opni/plugins/alerting/apis/alertops"
	metricsExporter "github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/metrics"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/proxy"
	"github.com/open-panoptes/opni/plugins/metrics/apis/cortexadmin"
	"github.com/open-panoptes/opni/plugins/metrics/apis/cortexops"
	metricsdk "go.opentelemetry.io/otel/sdk/metric"

	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/open-panoptes/opni/pkg/alerting/client"
	"github.com/open-panoptes/opni/pkg/alerting/server"
	"github.com/open-panoptes/opni/pkg/alerting/storage/spec"
	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
	httpext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/http"
	"github.com/open-panoptes/opni/pkg/plugins/apis/metrics"
	"github.com/open-panoptes/opni/pkg/util"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/alarms/v1"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/endpoints/v1"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/notifications/v1"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/node_backend"

	managementv1 "github.com/open-panoptes/opni/pkg/apis/management/v1"
	"github.com/open-panoptes/opni/pkg/logger"
	managementext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/management"
	streamext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/stream"
	"github.com/open-panoptes/opni/pkg/plugins/apis/capability"
	"github.com/open-panoptes/opni/pkg/plugins/apis/system"
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/pkg/util/future"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/drivers"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/apis/node"
)

func (p *Plugin) Components() []server.ServerComponent {
	return []server.ServerComponent{
		p.NotificationServerComponent,
		p.EndpointServerComponent,
		p.AlarmServerComponent,
		p.httpProxy,
	}
}

type Plugin struct {
	alertops.UnsafeAlertingAdminServer
	alertops.ConfigReconcilerServer
	system.UnimplementedSystemPluginClient

	ctx    context.Context
	logger *slog.Logger

	storageClientSet future.Future[spec.AlertingClientSet]

	alertingClient  future.Future[client.AlertingClient]
	clusterNotifier chan client.AlertingClient
	clusterDriver   future.Future[drivers.ClusterDriver]
	syncController  SyncController

	mgmtClient          future.Future[managementv1.ManagementClient]
	storageBackend      future.Future[storage.Backend]
	capabilitySpecStore future.Future[node_backend.CapabilitySpecKV]
	delegate            future.Future[streamext.StreamDelegate[agent.ClientSet]]
	adminClient         future.Future[cortexadmin.CortexAdminClient]
	cortexOpsClient     future.Future[cortexops.CortexOpsClient]
	natsConn            future.Future[*nats.Conn]
	js                  future.Future[nats.JetStreamContext]
	globalWatchers      management.ConditionWatcher

	gatewayConfig     future.Future[*v1beta1.GatewayConfig]
	alertingTLSConfig future.Future[*tls.Config]

	collector.CollectorServer

	*notifications.NotificationServerComponent
	*endpoints.EndpointServerComponent
	*alarms.AlarmServerComponent

	node node_backend.AlertingNodeBackend

	httpProxy *proxy.ProxyServer
	hsServer  *healthStatusServer
}

var (
	_ alertingv1.AlertEndpointsServer     = (*Plugin)(nil)
	_ alertingv1.AlertConditionsServer    = (*Plugin)(nil)
	_ alertingv1.AlertNotificationsServer = (*Plugin)(nil)
)

func NewPlugin(ctx context.Context) *Plugin {
	lg := logger.NewPluginLogger().WithGroup("alerting")
	storageClientSet := future.New[spec.AlertingClientSet]()
	metricReader := metricsdk.NewManualReader()
	metricsExporter.RegisterMeterProvider(metricsdk.NewMeterProvider(
		metricsdk.WithReader(metricReader),
	))
	collector := collector.NewCollectorServer(metricReader)
	p := &Plugin{
		ctx:    ctx,
		logger: lg,

		storageClientSet: storageClientSet,

		clusterNotifier: make(chan client.AlertingClient),
		clusterDriver:   future.New[drivers.ClusterDriver](),

		mgmtClient:          future.New[managementv1.ManagementClient](),
		storageBackend:      future.New[storage.Backend](),
		capabilitySpecStore: future.New[node_backend.CapabilitySpecKV](),
		delegate:            future.New[streamext.StreamDelegate[agent.ClientSet]](),

		adminClient:     future.New[cortexadmin.CortexAdminClient](),
		cortexOpsClient: future.New[cortexops.CortexOpsClient](),
		natsConn:        future.New[*nats.Conn](),
		js:              future.New[nats.JetStreamContext](),

		gatewayConfig:     future.New[*v1beta1.GatewayConfig](),
		alertingTLSConfig: future.New[*tls.Config](),
		alertingClient:    future.New[client.AlertingClient](),

		CollectorServer: collector,
	}

	p.syncController = NewSyncController(p.logger.With("component", "sync-controller"))
	p.hsServer = newHealthStatusServer(
		p.ready,
		p.healthy,
	)
	p.httpProxy = proxy.NewProxyServer(
		lg.With("component", "http-proxy"),
	)

	p.node = *node_backend.NewAlertingNodeBackend(
		p.logger.With("component", "node-backend"),
	)
	p.NotificationServerComponent = notifications.NewNotificationServerComponent(
		p.logger.With("component", "notifications"),
	)
	p.EndpointServerComponent = endpoints.NewEndpointServerComponent(
		p.ctx,
		p.logger.With("component", "endpoints"),
		p.NotificationServerComponent,
	)
	p.AlarmServerComponent = alarms.NewAlarmServerComponent(
		p.ctx,
		p.logger.With("component", "alarms"),
		p.NotificationServerComponent,
	)

	future.Wait4(
		p.storageBackend,
		p.mgmtClient,
		p.capabilitySpecStore,
		p.delegate,
		func(
			storageBackend storage.Backend,
			mgmtClient managementv1.ManagementClient,
			specStore node_backend.CapabilitySpecKV,
			delegate streamext.StreamDelegate[agent.ClientSet],
		) {
			p.node.Initialize(specStore, mgmtClient, delegate, storageBackend)
		},
	)

	future.Wait1(
		p.alertingTLSConfig,
		func(tlsConfig *tls.Config) {
			p.httpProxy.Initialize(tlsConfig)

			alertingClient, err := client.NewClient(
				client.WithTLSConfig(tlsConfig),
			)
			if err != nil {
				panic(err)
			}
			p.alertingClient.Set(alertingClient)
		},
	)

	future.Wait2(p.storageClientSet, p.alertingClient, func(s spec.AlertingClientSet, alertingClient client.AlertingClient) {
		serverCfg := server.Config{
			Client: alertingClient.Clone(),
		}
		p.NotificationServerComponent.SetConfig(
			serverCfg,
		)

		p.EndpointServerComponent.SetConfig(
			serverCfg,
		)
		p.NotificationServerComponent.Initialize(notifications.NotificationServerConfiguration{
			ConditionStorage: s.Conditions(),
			EndpointStorage:  s.Endpoints(),
		})

		p.EndpointServerComponent.Initialize(endpoints.EndpointServerConfiguration{
			ConditionStorage: s.Conditions(),
			EndpointStorage:  s.Endpoints(),
			RouterStorage:    s.Routers(),
			HashRing:         s,
		})

	})

	future.Wait6(
		p.js, p.storageClientSet, p.mgmtClient, p.adminClient, p.cortexOpsClient, p.alertingClient,
		func(
			js nats.JetStreamContext,
			s spec.AlertingClientSet,
			mgmtClient managementv1.ManagementClient,
			adminClient cortexadmin.CortexAdminClient,
			cortexOpsClient cortexops.CortexOpsClient,
			alertingClient client.AlertingClient,
		) {
			serverCfg := server.Config{
				Client: alertingClient.Clone(),
			}

			p.AlarmServerComponent.SetConfig(
				serverCfg,
			)
			p.AlarmServerComponent.Initialize(alarms.AlarmServerConfiguration{
				ConditionStorage: s.Conditions(),
				IncidentStorage:  s.Incidents(),
				StateStorage:     s.States(),
				RouterStorage:    s.Routers(),
				MgmtClient:       mgmtClient,
				AdminClient:      adminClient,
				CortexOpsClient:  cortexOpsClient,
				Js:               js,
			})
		})
	return p
}

var (
	_ alertingv1.AlertEndpointsServer     = (*Plugin)(nil)
	_ alertingv1.AlertConditionsServer    = (*Plugin)(nil)
	_ alertingv1.AlertNotificationsServer = (*Plugin)(nil)
)

func Scheme(ctx context.Context) meta.Scheme {
	scheme := meta.NewScheme()
	p := NewPlugin(ctx)
	scheme.Add(system.SystemPluginID, system.NewPlugin(p))
	scheme.Add(httpext.HTTPAPIExtensionPluginID, httpext.NewPlugin(p))
	scheme.Add(managementext.ManagementAPIExtensionPluginID,
		managementext.NewPlugin(
			util.PackService(
				&alertingv1.AlertConditions_ServiceDesc,
				p.AlarmServerComponent,
			),
			util.PackService(
				&alertingv1.AlertEndpoints_ServiceDesc,
				p.EndpointServerComponent,
			),
			util.PackService(
				&alertingv1.AlertNotifications_ServiceDesc,
				p.NotificationServerComponent,
			),
			util.PackService(
				&alertops.AlertingAdmin_ServiceDesc,
				p,
			),
			util.PackService(
				&alertops.ConfigReconciler_ServiceDesc,
				p,
			),
			util.PackService(
				&node.AlertingNodeConfiguration_ServiceDesc,
				&p.node,
			),
			util.PackService(
				&node.NodeAlertingCapability_ServiceDesc,
				&p.node,
			),
		),
	)

	scheme.Add(metrics.MetricsPluginID, metrics.NewPlugin(p))
	scheme.Add(capability.CapabilityBackendPluginID, capability.NewPlugin(&p.node))
	scheme.Add(streamext.StreamAPIExtensionPluginID, streamext.NewGatewayPlugin(p))
	return scheme
}
