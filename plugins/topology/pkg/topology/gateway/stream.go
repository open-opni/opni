package gateway

import (
	"github.com/open-panoptes/opni/pkg/agent"
	"github.com/open-panoptes/opni/pkg/capabilities/wellknown"
	streamext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/stream"
	"github.com/open-panoptes/opni/plugins/topology/apis/node"
	"github.com/open-panoptes/opni/plugins/topology/apis/stream"
	"google.golang.org/grpc"
)

func (p *Plugin) StreamServers() []streamext.Server {
	return []streamext.Server{
		{
			Desc:              &stream.RemoteTopology_ServiceDesc,
			Impl:              &p.topologyRemoteWrite,
			RequireCapability: wellknown.CapabilityTopology,
		},
		{
			Desc:              &node.NodeTopologyCapability_ServiceDesc,
			Impl:              &p.topologyBackend,
			RequireCapability: wellknown.CapabilityTopology,
		},
	}
}

func (p *Plugin) UseStreamClient(cc grpc.ClientConnInterface) {
	p.delegate.Set(streamext.NewDelegate(cc, agent.NewClientSet))
}
