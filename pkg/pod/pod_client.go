package pod

import (
	v1 "k8s.io/api/core/v1"
	metricsType "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type podClient interface {
	GetPods(options string) (*v1.PodList, error)
	GetPod(name string) (*v1.Pod, error)
	DeletePod(name string) error
	GetPodMetrics(name string) (*metricsType.PodMetrics, error)
	GetPodsMetrics(options string) (*metricsType.PodMetricsList, error)
}
