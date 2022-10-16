package configs

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type KubernetesConfig struct {
	MetricsNamespace string
	Namespace        string
	ClientSet        *kubernetes.Clientset
	Metrics          *metrics.Clientset
}

func NewKubernetesConfig(namespace, metricsNamespace string) (*KubernetesConfig, error) {
	K8sConfig := &KubernetesConfig{
		MetricsNamespace: metricsNamespace,
		Namespace:        namespace,
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		return K8sConfig, err
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return K8sConfig, err
	}

	metricsv, err := metrics.NewForConfig(config)
	if err != nil {
		return K8sConfig, err
	}

	K8sConfig.ClientSet = client
	K8sConfig.Metrics = metricsv

	return K8sConfig, nil
}
