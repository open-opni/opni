package drivers

import (
	"context"

	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	"github.com/open-panoptes/opni/pkg/plugins/driverutil"
	"github.com/open-panoptes/opni/plugins/metrics/apis/cortexops"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ClusterDriver interface {
	cortexops.SpecializedConfigServer
	cortexops.SpecializedDryRunServer
	driverutil.InstallerServer

	// ShouldDisableNode is called during node sync for nodes which otherwise
	// have this capability enabled. If this function returns an error, the
	// node will be set to disabled instead, and the error will be logged.
	ShouldDisableNode(*corev1.Reference) error
	ListPresets(context.Context, *emptypb.Empty) (*cortexops.PresetList, error)
}

var ClusterDrivers = driverutil.NewDriverCache[ClusterDriver]()

type NoopClusterDriver struct {
	cortexops.UnimplementedCortexOpsServer
}

func (d *NoopClusterDriver) Name() string {
	return "noop"
}

func (d *NoopClusterDriver) ShouldDisableNode(*corev1.Reference) error {
	// the noop driver will never forcefully disable a node
	return nil
}

func init() {
	ClusterDrivers.Register("noop", func(context.Context, ...driverutil.Option) (ClusterDriver, error) {
		return &NoopClusterDriver{}, nil
	})
}
