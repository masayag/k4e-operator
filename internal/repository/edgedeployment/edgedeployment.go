package edgedeployment

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/jakub-dzon/k4e-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//go:generate mockgen -package=edgedeployment -destination=mock_edgedeployment.go . Repository
type Repository interface {
	Read(ctx context.Context, name string, namespace string) (*v1alpha1.EdgeDeployment, error)
	Patch(ctx context.Context, old, new *v1alpha1.EdgeDeployment) error
	RemoveFinalizer(ctx context.Context, edgeDeployment *v1alpha1.EdgeDeployment, finalizer string) error
	ListByLabel(ctx context.Context, labelName, labelValue string) ([]v1alpha1.EdgeDeployment, error)
}

type CRRespository struct {
	client client.Client
}

func NewEdgeDeploymentRepository(client client.Client) *CRRespository {
	return &CRRespository{client: client}
}

func (r *CRRespository) Read(ctx context.Context, name string, namespace string) (*v1alpha1.EdgeDeployment, error) {
	edgeDeployment := v1alpha1.EdgeDeployment{}
	err := r.client.Get(ctx, client.ObjectKey{Namespace: namespace, Name: name}, &edgeDeployment)
	return &edgeDeployment, err
}

func (r *CRRespository) Patch(ctx context.Context, old, new *v1alpha1.EdgeDeployment) error {
	patch := client.MergeFrom(old)
	return r.client.Patch(ctx, new, patch)
}

func (r *CRRespository) RemoveFinalizer(ctx context.Context, edgeDeployment *v1alpha1.EdgeDeployment, finalizer string) error {
	cp := edgeDeployment.DeepCopy()

	var finalizers []string
	for _, f := range cp.Finalizers {
		if f != finalizer {
			finalizers = append(finalizers, f)
		}
	}
	cp.Finalizers = finalizers

	err := r.Patch(ctx, edgeDeployment, cp)
	if err == nil {
		edgeDeployment.Finalizers = cp.Finalizers
	}

	return nil
}

func (r *CRRespository) ListByLabel(ctx context.Context, labelName, labelValue string) ([]v1alpha1.EdgeDeployment, error) {
	edgeDeployments := v1alpha1.EdgeDeploymentList{}
	requirement, err := labels.NewRequirement(labelName, selection.Equals, []string{labelValue})
	if err != nil {
		return nil, err
	}
	selector := labels.NewSelector().Add(*requirement)
	err = r.client.List(ctx, &edgeDeployments, client.MatchingLabelsSelector{Selector: selector})
	if err != nil {
		return nil, err
	}
	return edgeDeployments.Items, nil
}
