package deployment

import "service-client/pkg/pod"

type Deployment struct {
	Name           string    `json:"name"`
	ID             string    `json:"id"`
	NumberOfPods   int32     `json:"numberOfPods"`
	ContainerImage string    `json:"containerImage"`
	Pods           []pod.Pod `json:"pods"`
}

type DeploymentUpdateRequest struct {
	Name           string `json:"-"`
	Replicas       int32  `json:"replicas,omitempty"`
	ContainerImage string `json:"containerImage,omitempty"`
}
