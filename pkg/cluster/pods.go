package cluster

import (
	"context"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	metricsType "k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metricsv1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"
)

type K8SPodClient struct {
	Context   context.Context
	Interface corev1.PodInterface
	Metrics   metricsv1.PodMetricsInterface
}

func (k K8SPodClient) GetPods(options string) (*core.PodList, error) {
	return k.Interface.List(k.Context, metav1.ListOptions{LabelSelector: options})
}

func (k K8SPodClient) GetPod(name string) (*core.Pod, error) {
	return k.Interface.Get(k.Context, name, metav1.GetOptions{})
}

func (k K8SPodClient) DeletePod(name string) error {
	return k.Interface.Delete(k.Context, name, metav1.DeleteOptions{})
}

func (k K8SPodClient) GetPodMetrics(name string) (*metricsType.PodMetrics, error) {
	return k.Metrics.Get(k.Context, name, metav1.GetOptions{})
}

func (k K8SPodClient) GetPodsMetrics(options string) (*metricsType.PodMetricsList, error) {
	return k.Metrics.List(k.Context, metav1.ListOptions{})
}
