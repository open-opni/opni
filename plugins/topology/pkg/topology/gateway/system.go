package gateway

import (
	"context"
	"os"

	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	managementv1 "github.com/open-panoptes/opni/pkg/apis/management/v1"
	"github.com/open-panoptes/opni/pkg/config/v1beta1"
	"github.com/open-panoptes/opni/pkg/logger"
	"github.com/open-panoptes/opni/pkg/machinery"
	"github.com/open-panoptes/opni/pkg/plugins/apis/system"
	"github.com/open-panoptes/opni/pkg/task"
	natsutil "github.com/open-panoptes/opni/pkg/util/nats"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	_ "github.com/open-panoptes/opni/pkg/storage/etcd"
	_ "github.com/open-panoptes/opni/pkg/storage/jetstream"
)

func (p *Plugin) UseManagementAPI(client managementv1.ManagementClient) {
	p.mgmtClient.Set(client)
	cfg, err := client.GetConfig(
		context.Background(),
		&emptypb.Empty{},
		grpc.WaitForReady(true),
	)
	if err != nil {
		p.logger.With(logger.Err(err)).Error("failed to get config")
		os.Exit(1)
	}

	objectList, err := machinery.LoadDocuments(cfg.Documents)
	if err != nil {
		p.logger.With(logger.Err(err)).Error("failed to load config")
		os.Exit(1)
	}
	machinery.LoadAuthProviders(p.ctx, objectList)
	objectList.Visit(func(config *v1beta1.GatewayConfig) {
		backend, err := machinery.ConfigureStorageBackend(p.ctx, &config.Spec.Storage)
		if err != nil {
			p.logger.With(
				"err", err,
			).Error("failed to configure storage backend")
			os.Exit(1)
		}
		p.storageBackend.Set(backend)
		p.gatewayConfig.Set(config)
		p.configureTopologyManagement()
	})
	<-p.ctx.Done()
}

func (p *Plugin) UseKeyValueStore(client system.KeyValueStoreClient) {
	// set other futures before trying to acquire NATS connection
	ctrl, err := task.NewController(
		p.ctx,
		"topology.uninstall",
		system.NewKVStoreClient[*corev1.TaskStatus](client),
		&p.uninstallRunner)

	if err != nil {
		p.logger.With(
			logger.Err(err),
		).Error("failed to create uninstall task controller")
	}
	p.uninstallController.Set(ctrl)

	p.storage.Set(ConfigStorageAPIs{
		Placeholder: system.NewKVStoreClient[proto.Message](client),
	})
	cfg := p.gatewayConfig.Get().Spec.Storage.JetStream
	natsURL := os.Getenv("NATS_SERVER_URL")
	natsSeedPath := os.Getenv("NKEY_SEED_FILENAME")
	if cfg == nil {
		cfg = &v1beta1.JetStreamStorageSpec{}
	}
	if cfg.Endpoint == "" {
		cfg.Endpoint = natsURL
	}
	if cfg.NkeySeedPath == "" {
		cfg.NkeySeedPath = natsSeedPath
	}
	nc, err := natsutil.AcquireNATSConnection(p.ctx, cfg)
	if err != nil {
		p.logger.With(
			logger.Err(err),
		).Error("fatal :  failed to acquire NATS connection")
		os.Exit(1)
	}
	p.nc.Set(nc)
	<-p.ctx.Done()
}
