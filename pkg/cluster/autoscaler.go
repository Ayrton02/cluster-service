package cluster

import (
	"context"

	v2 "k8s.io/api/autoscaling/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	autoscalingv2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2"
)

type K8SAutoscalerClient struct {
	Context   context.Context
	Interface autoscalingv2.HorizontalPodAutoscalerInterface
}

func (k K8SAutoscalerClient) GetAutoscaler(name string) (*v2.HorizontalPodAutoscaler, error) {
	return k.Interface.Get(k.Context, name, metav1.GetOptions{})
}

func (k K8SAutoscalerClient) GetAutoscalers() (*v2.HorizontalPodAutoscalerList, error) {
	return k.Interface.List(k.Context, metav1.ListOptions{})
}

func (k K8SAutoscalerClient) UpdateAutoscaler(autoscaler *v2.HorizontalPodAutoscaler) (*v2.HorizontalPodAutoscaler, error) {
	return k.Interface.Update(k.Context, autoscaler, metav1.UpdateOptions{})
}

func (k K8SAutoscalerClient) CreateAutoscaler(autoscaler *v2.HorizontalPodAutoscaler) (*v2.HorizontalPodAutoscaler, error) {
	return k.Interface.Create(k.Context, autoscaler, metav1.CreateOptions{})
}
