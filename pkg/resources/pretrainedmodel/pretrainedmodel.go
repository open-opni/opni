package pretrainedmodel

import (
	"context"
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"

	"emperror.dev/errors"
	"github.com/cisco-open/operator-tools/pkg/reconciler"
	aiv1beta1 "github.com/open-panoptes/opni/apis/ai/v1beta1"
	"github.com/open-panoptes/opni/pkg/resources"
	"github.com/open-panoptes/opni/pkg/resources/hyperparameters"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type Reconciler struct {
	reconciler.ResourceReconciler
	ctx    context.Context
	client client.Client
	model  *aiv1beta1.PretrainedModel
}

func NewReconciler(
	ctx context.Context,
	client client.Client,
	instance *aiv1beta1.PretrainedModel,
	opts ...reconciler.ResourceReconcilerOption,
) *Reconciler {
	return &Reconciler{
		ResourceReconciler: reconciler.NewReconcilerWith(client,
			append(opts, reconciler.WithLog(log.FromContext(ctx)))...),
		ctx:    ctx,
		model:  instance,
		client: client,
	}
}

func (r *Reconciler) hyperparameters() (runtime.Object, reconciler.DesiredState, error) {
	cm, err := hyperparameters.GenerateHyperparametersConfigMap(r.model.Name, r.model.Namespace, r.model.Spec.Hyperparameters)
	if err != nil {
		return nil, nil, err
	}
	cm.Labels[resources.PretrainedModelLabel] = r.model.Name
	ctrl.SetControllerReference(r.model, &cm, r.client.Scheme())
	return &cm, reconciler.StatePresent, err
}

func (r *Reconciler) Reconcile() (*reconcile.Result, error) {
	for _, factory := range []resources.Resource{
		r.hyperparameters,
	} {
		o, state, err := factory()
		if err != nil {
			return nil, errors.WrapIf(err, "failed to create object")
		}
		if o == nil {
			panic(fmt.Sprintf("reconciler %#v created a nil object", factory))
		}
		result, err := r.ReconcileResource(o, state)
		if err != nil {
			return nil, errors.WrapWithDetails(err, "failed to reconcile resource",
				"resource", o.GetObjectKind().GroupVersionKind())
		}
		if result != nil {
			return result, nil
		}
	}
	return nil, nil
}
