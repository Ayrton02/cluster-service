package cluster

import (
	"context"

	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type K8SClient struct {
	Context     context.Context
	Pods        *K8SPodClient
	Autoscalers *K8SAutoscalerClient
	Deployment  *K8SDeploymentClient
	Node        *K8SNodeClient
}

type K8SClientSets struct {
	Client  *kubernetes.Clientset
	Metrics metrics.Clientset
}

type K8SNamespaces struct {
	Operational string
	Metrics     string
}

func NewK8SClient(Ctx context.Context, sets K8SClientSets, namespaces K8SNamespaces) *K8SClient {
	return &K8SClient{
		Context: Ctx,
		Pods: &K8SPodClient{
			Context:   Ctx,
			Interface: sets.Client.CoreV1().Pods(namespaces.Operational),
			Metrics:   sets.Metrics.MetricsV1beta1().PodMetricses(namespaces.Operational),
		},
		Autoscalers: &K8SAutoscalerClient{
			Context:   Ctx,
			Interface: sets.Client.AutoscalingV2().HorizontalPodAutoscalers(namespaces.Operational),
		},
		Deployment: &K8SDeploymentClient{
			Context:   Ctx,
			Interface: sets.Client.AppsV1().Deployments(namespaces.Operational),
		},
		Node: &K8SNodeClient{
			Context:   Ctx,
			Interface: sets.Client.CoreV1().Nodes(),
			Metrics:   sets.Metrics.MetricsV1beta1().NodeMetricses(),
		},
	}
}
