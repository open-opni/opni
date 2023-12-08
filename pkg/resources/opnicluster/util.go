package opnicluster

import (
	aiv1beta1 "github.com/open-panoptes/opni/apis/ai/v1beta1"
	"github.com/open-panoptes/opni/pkg/resources"
	opnimeta "github.com/open-panoptes/opni/pkg/util/meta"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func (r *Reconciler) serviceLabels(service aiv1beta1.ServiceKind) map[string]string {
	return map[string]string{
		resources.AppNameLabel: service.ServiceName(),
		resources.ServiceLabel: service.String(),
		resources.PartOfLabel:  "opni",
	}
}

func (r *Reconciler) natsLabels() map[string]string {
	return map[string]string{
		resources.AppNameLabel:    "nats",
		resources.PartOfLabel:     "opni",
		resources.OpniClusterName: r.opniCluster.Name,
	}
}

func (r *Reconciler) pretrainedModelLabels(modelName string) map[string]string {
	return map[string]string{
		resources.PretrainedModelLabel: modelName,
	}
}

func (r *Reconciler) serviceImageSpec(service aiv1beta1.ServiceKind) opnimeta.ImageSpec {
	return opnimeta.ImageResolver{
		Version:             r.opniCluster.Spec.Version,
		ImageName:           service.ImageName(),
		DefaultRepo:         "docker.io/rancher",
		DefaultRepoOverride: r.opniCluster.Spec.DefaultRepo,
		ImageOverride:       service.GetImageSpec(r.opniCluster),
	}.Resolve()
}

func (r *Reconciler) serviceNodeSelector(service aiv1beta1.ServiceKind) map[string]string {
	if s := service.GetNodeSelector(r.opniCluster); len(s) > 0 {
		return s
	}
	return r.opniCluster.Spec.GlobalNodeSelector
}

func (r *Reconciler) serviceTolerations(service aiv1beta1.ServiceKind) []corev1.Toleration {
	return append(r.opniCluster.Spec.GlobalTolerations, service.GetTolerations(r.opniCluster)...)
}

func addCPUInferenceLabel(deployment *appsv1.Deployment) {
	deployment.Labels[resources.OpniInferenceType] = "cpu"
	deployment.Spec.Template.Labels[resources.OpniInferenceType] = "cpu"
	deployment.Spec.Selector.MatchLabels[resources.OpniInferenceType] = "cpu"
}
