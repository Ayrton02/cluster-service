package cluster

import (
	"context"

	apps "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type K8SDeploymentClient struct {
	Context   context.Context
	Interface appsv1.DeploymentInterface
}

func (k K8SDeploymentClient) GetDeployments() (*apps.DeploymentList, error) {
	return k.Interface.List(k.Context, metav1.ListOptions{})
}

func (k K8SDeploymentClient) GetDeployment(name string) (*apps.Deployment, error) {
	return k.Interface.Get(k.Context, name, metav1.GetOptions{})
}

func (k K8SDeploymentClient) DeleteDeployment(name string) error {
	return k.Interface.Delete(k.Context, name, metav1.DeleteOptions{})
}

func (k K8SDeploymentClient) UpdateDeployment(deployment *apps.Deployment) (*apps.Deployment, error) {
	return k.Interface.Update(k.Context, deployment, metav1.UpdateOptions{})
}

func (k K8SDeploymentClient) CreateDeployment(deployment *apps.Deployment) (*apps.Deployment, error) {
	return k.Interface.Create(k.Context, deployment, metav1.CreateOptions{})
}
