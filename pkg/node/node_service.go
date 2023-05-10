package node

import (
	"github.com/gorilla/mux"
)

var Service *nodeService

type nodeService struct {
	client nodeClient
}

func InitNodeService(client nodeClient, route *mux.Router) {
	registerNodeRoutes(route)
	Service = &nodeService{
		client: client,
	}
}

func (s nodeService) GetNodes() ([]Node, error) {
	var nodes []Node
	res, err := s.client.GetNodes()
	if err != nil {
		return nodes, err
	}

	for _, i := range res.Items {

		nodes = append(nodes, Node{
			Name:            i.Name,
			UUID:            string(i.UID),
			MemoryAllocated: i.Status.Allocatable.Memory().String(),
			CPUAllocated:    i.Status.Allocatable.Cpu().String(),
		})
	}

	return nodes, err
}

func (s nodeService) GetSingleNode(name string) (Node, error) {
	res, err := s.client.GetNode(name)
	if err != nil {
		return Node{}, err
	}

	return Node{
		Name:            res.Name,
		UUID:            string(res.UID),
		MemoryAllocated: res.Status.Allocatable.Memory().String(),
		CPUAllocated:    res.Status.Allocatable.Cpu().String(),
	}, err

}

func (s nodeService) GetSingleNodeMetrics(name string) (NodeMetrics, error) {
	res, err := s.client.GetNodeMetrics(name)
	if err != nil {
		return NodeMetrics{}, err
	}

	return NodeMetrics{
		Name:        res.Name,
		MemoryUsage: res.Usage.Memory().String(),
		CPUUsage:    res.Usage.Cpu().String(),
		UUID:        string(res.ObjectMeta.UID),
	}, err

}

func (s nodeService) GetNodesMetrics() ([]NodeMetrics, error) {
	var metrics []NodeMetrics
	res, err := s.client.GetNodesMetrics()
	if err != nil {
		return metrics, err
	}

	for _, i := range res.Items {

		metrics = append(metrics,
			NodeMetrics{
				Name:        i.Name,
				MemoryUsage: i.Usage.Memory().String(),
				CPUUsage:    i.Usage.Cpu().String(),
				UUID:        string(i.ObjectMeta.UID),
			})
	}

	return metrics, err
}
