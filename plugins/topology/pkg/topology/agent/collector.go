package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"log/slog"

	controlv1 "github.com/open-panoptes/opni/pkg/apis/control/v1"
	"github.com/open-panoptes/opni/pkg/health"
	"github.com/open-panoptes/opni/pkg/logger"
	"github.com/open-panoptes/opni/pkg/topology/graph"
	"github.com/open-panoptes/opni/plugins/topology/apis/node"
	"github.com/open-panoptes/opni/plugins/topology/apis/stream"
	"google.golang.org/protobuf/types/known/emptypb"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type BatchingConfig struct {
	maxSize int
	timeout time.Duration
}

type TopologyStreamer struct {
	logger     *slog.Logger
	conditions health.ConditionTracker

	v                chan client.Object
	eventWatchClient client.WithWatch

	identityClientMu       sync.Mutex
	identityClient         controlv1.IdentityClient
	topologyStreamClientMu sync.Mutex
	topologyStreamClient   stream.RemoteTopologyClient
}

func NewTopologyStreamer(ct health.ConditionTracker, lg *slog.Logger) *TopologyStreamer {
	return &TopologyStreamer{
		// FIXME: reintroduce this when we want to monitor kubernetes events
		// eventWatchClient: util.Must(client.NewWithWatch(
		// 	util.Must(rest.InClusterConfig()),
		// 	client.Options{
		// 		Scheme: apis.NewScheme(),
		// 	})),
		logger:     lg,
		conditions: ct,
	}
}

func (s *TopologyStreamer) SetTopologyStreamClient(client stream.RemoteTopologyClient) {
	s.topologyStreamClientMu.Lock()
	defer s.topologyStreamClientMu.Unlock()
	s.topologyStreamClient = client
}

func (s *TopologyStreamer) SetIdentityClient(identityClient controlv1.IdentityClient) {
	s.identityClientMu.Lock()
	defer s.identityClientMu.Unlock()
	s.identityClient = identityClient

}

func (s *TopologyStreamer) Run(ctx context.Context, spec *node.TopologyCapabilitySpec) error {
	lg := s.logger
	if spec == nil {
		lg.With("stream", "topology").Warn("no topology capability spec provided, setting defaults")

		// set some sensible defaults
	}
	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()

	// blocking action
	for {
		select {
		case <-ctx.Done():
			lg.With(
				logger.Err(ctx.Err()),
			).Warn("topology stream closing")
			return nil
		case <-tick.C:
			// will panic if not in a cluster
			g, err := graph.TraverseTopology(lg, graph.NewRuntimeFactory())
			if err != nil {
				lg.With(
					logger.Err(err),
				).Error("Could not construct topology graph")
			}
			var b bytes.Buffer
			err = json.NewEncoder(&b).Encode(g)
			if err != nil {
				lg.With(
					logger.Err(err),
				).Warn("failed to encode kubernetes graph")
				continue
			}
			s.identityClientMu.Lock()
			thisCluster, err := s.identityClient.Whoami(ctx, &emptypb.Empty{})
			if err != nil {
				lg.With(
					logger.Err(err),
				).Warn("failed to get cluster identity")
				continue
			}
			s.identityClientMu.Unlock()

			s.topologyStreamClientMu.Lock()
			_, err = s.topologyStreamClient.Push(ctx, &stream.Payload{
				Graph: &stream.TopologyGraph{
					ClusterId: thisCluster,
					Data:      b.Bytes(),
					Repr:      stream.GraphRepr_KubectlGraph,
				},
			})
			if err != nil {
				lg.Error(fmt.Sprintf("failed to push topology graph: %s", err))
			}
			s.topologyStreamClientMu.Unlock()
		}
	}
}
