package autoscaler

type PodAutoscaler struct {
	Name      string         `json:"name"`
	ID        string         `json:"id"`
	Replicas  PodReplicas    `json:"replicas"`
	Resources []PodResources `json:"resources"`
}

type PodReplicas struct {
	Min int32 `json:"min"`
	Max int32 `json:"max"`
}

type PodResources struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PodAutoscalerUpdateRequest struct {
	Name      string                `json:"-"`
	Replicas  PodReplicas           `json:"replicas"`
	Resources []PodResourcesRequest `json:"resources"`
}

type PodAutoscalerCreateRequest struct {
	Name             string                `json:"name"`
	Replicas         PodReplicas           `json:"replicas"`
	Resources        []PodResourcesRequest `json:"resources"`
	DeploymentTarget string                `json:"deploymentTarget"`
}

type PodResourcesRequest struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}
