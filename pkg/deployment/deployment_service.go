package deployment

import (
	"fmt"
	"service-client/pkg/pod"

	"github.com/gorilla/mux"
	v1 "k8s.io/api/apps/v1"
)

var Service *deploymentService

type deploymentService struct {
	client deploymentClient
}

func InitDeploymentService(client deploymentClient, route *mux.Router) {
	registerDeploymentRoutes(route)
	Service = &deploymentService{
		client: client,
	}
}

func (s deploymentService) GetDeployments() ([]Deployment, error) {
	var deployments []Deployment

	res, err := s.client.GetDeployments()

	if err != nil {
		return deployments, err
	}

	for _, d := range res.Items {
		deployment, err := toDeployment(&d)
		if err != nil {
			return deployments, err
		}

		deployments = append(deployments, deployment)
	}

	return deployments, nil
}

func (s deploymentService) GetSingleDeployment(name string) (Deployment, error) {
	var deployment Deployment

	res, err := s.client.GetDeployment(name)

	if err != nil {
		return deployment, err
	}

	return toDeployment(res)
}

func (s deploymentService) DeleteSingleDeployment(name string) error {
	return s.client.DeleteDeployment(name)
}

func (s deploymentService) UpdateSingleDeployment(updateDeployment DeploymentUpdateRequest) error {
	k8sDeployment, err := s.client.GetDeployment(updateDeployment.Name)
	if err != nil {
		return err
	}

	prepareToUpdate(k8sDeployment, updateDeployment)

	_, err = s.client.UpdateDeployment(k8sDeployment)

	return err
}

func toDeployment(k8sDeployment *v1.Deployment) (Deployment, error) {
	var deployment Deployment
	pods := make([]pod.Pod, 0)
	labels := make([]string, 0, len(k8sDeployment.Labels))

	for k, v := range k8sDeployment.Labels {
		options := fmt.Sprintf("%s=%s", k, v)
		labels = append(labels, options)
	}

	for _, l := range labels {
		labeledPods, err := pod.Service.GetPods(l)
		if err != nil {
			return deployment, err
		}
		pods = append(pods, labeledPods...)
	}

	deployment = Deployment{
		Name:           k8sDeployment.Name,
		ID:             string(k8sDeployment.UID),
		Pods:           pods,
		ContainerImage: k8sDeployment.Spec.Template.Spec.Containers[0].Image,
		NumberOfPods:   *k8sDeployment.Spec.Replicas,
	}

	return deployment, nil
}

func prepareToUpdate(k8sDeployment *v1.Deployment, updateDeployment DeploymentUpdateRequest) {
	if updateDeployment.Replicas != 0 {
		k8sDeployment.Spec.Replicas = &updateDeployment.Replicas
	}

	if updateDeployment.ContainerImage != "" {
		for _, c := range k8sDeployment.Spec.Template.Spec.Containers {
			c.Image = updateDeployment.ContainerImage
		}
	}
}
