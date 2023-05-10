package node

import (
	core "k8s.io/api/core/v1"
	metricsType "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

type nodeClient interface {
	GetNodes() (*core.NodeList, error)
	GetNode(name string) (*core.Node, error)
	GetNodeMetrics(name string) (*metricsType.NodeMetrics, error)
}
