package gateway

import (
	controlv1 "github.com/open-panoptes/opni/pkg/apis/control/v1"
	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	"github.com/open-panoptes/opni/pkg/auth/cluster"
	"github.com/open-panoptes/opni/pkg/storage"
	"github.com/open-panoptes/opni/pkg/update"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewLastKnownDetailsApplier(storageBackend storage.ClusterStore) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		id := cluster.StreamAuthorizedID(ss.Context())

		lkcd := &corev1.LastKnownConnectionDetails{
			Time: timestamppb.Now(),
		}

		// best effort peer info
		if p, ok := peer.FromContext(ss.Context()); ok {
			lkcd.Address = p.Addr.String()
		}

		mmd, ok := update.ManifestMetadataFromContext(ss.Context())
		if ok {
			lkcd.PluginVersions = mmd.DigestMap()
		}

		md, ok := metadata.FromIncomingContext(ss.Context())
		if ok {
			values := md.Get(controlv1.AgentBuildInfoKey)
			if len(values) > 0 {
				buildInfo := &corev1.BuildInfo{}
				if err := protojson.Unmarshal([]byte(values[0]), buildInfo); err != nil {
					return err
				}

				lkcd.AgentBuildInfo = buildInfo
			}
		}

		if _, err := storageBackend.UpdateCluster(ss.Context(), &corev1.Reference{Id: id}, func(cluster *corev1.Cluster) {
			cluster.Metadata.LastKnownConnectionDetails = lkcd
		}); err != nil {
			return status.Errorf(codes.Internal, "failed to update cluster: %v", err)
		}
		return handler(srv, ss)
	}
}
