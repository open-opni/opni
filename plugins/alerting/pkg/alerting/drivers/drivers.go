package drivers

import (
	"context"

	"github.com/open-panoptes/opni/pkg/alerting/drivers/config"
	"github.com/open-panoptes/opni/pkg/alerting/shared"
	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	"github.com/open-panoptes/opni/pkg/plugins/driverutil"
	"github.com/open-panoptes/opni/plugins/alerting/apis/alertops"
)

type ClusterDriver interface {
	alertops.AlertingAdminServer
	// ShouldDisableNode is called during node sync for nodes which otherwise
	// have this capability enabled. If this function returns an error, the
	// node will be set to disabled instead, and the error will be logged.
	ShouldDisableNode(*corev1.Reference) error
	GetDefaultReceiver() *config.WebhookConfig
}

var Drivers = driverutil.NewDriverCache[ClusterDriver]()

type NoopClusterDriver struct {
	alertops.UnimplementedAlertingAdminServer
}

func (d *NoopClusterDriver) ShouldDisableNode(*corev1.Reference) error {
	// the noop driver will never forcefully disable a node
	return nil
}

func (d *NoopClusterDriver) GetRuntimeOptions() shared.AlertingClusterOptions {
	return shared.AlertingClusterOptions{}
}

func (d *NoopClusterDriver) GetDefaultReceiver() *config.WebhookConfig {
	return &config.WebhookConfig{}
}

func init() {
	Drivers.Register("noop", func(ctx context.Context, opts ...driverutil.Option) (ClusterDriver, error) {
		return &NoopClusterDriver{}, nil
	})
}
