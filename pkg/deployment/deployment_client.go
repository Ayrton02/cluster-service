package deployment

import apps "k8s.io/api/apps/v1"

type deploymentClient interface {
	GetDeployments() (*apps.DeploymentList, error)
	GetDeployment(name string) (*apps.Deployment, error)
	DeleteDeployment(name string) error
	UpdateDeployment(deployment *apps.Deployment) (*apps.Deployment, error)
	CreateDeployment(deployment *apps.Deployment) (*apps.Deployment, error)
}
