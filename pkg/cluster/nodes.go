package cluster

import (
	"context"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	metricsType "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsv1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
)

type K8SNodeClient struct {
	Context   context.Context
	Interface corev1.NodeInterface
	Metrics   metricsv1.NodeMetricsInterface
}

func (k K8SNodeClient) GetNodes() (*core.NodeList, error) {
	return k.Interface.List(k.Context, metav1.ListOptions{})
}

func (k K8SNodeClient) GetNode(name string) (*core.Node, error) {
	return k.Interface.Get(k.Context, name, metav1.GetOptions{})
}

func (k K8SNodeClient) GetNodeMetrics(name string) (*metricsType.NodeMetrics, error) {
	return k.Metrics.Get(k.Context, name, metav1.GetOptions{})
}
