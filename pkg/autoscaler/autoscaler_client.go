package autoscaler

import v2 "k8s.io/api/autoscaling/v2"

type autoscalerClient interface {
	GetAutoscaler(name string) (*v2.HorizontalPodAutoscaler, error)
	GetAutoscalers() (*v2.HorizontalPodAutoscalerList, error)
	UpdateAutoscaler(autoscaler *v2.HorizontalPodAutoscaler) (*v2.HorizontalPodAutoscaler, error)
	CreateAutoscaler(autoscaler *v2.HorizontalPodAutoscaler) (*v2.HorizontalPodAutoscaler, error)
}
