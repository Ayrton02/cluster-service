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
			Name:   i.Name,
			UUID:   string(i.UID),
			Memory: i.Status.Allocatable.Memory().String(),
			CPU:    i.Status.Allocatable.Cpu().String(),
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
		Name:   res.Name,
		UUID:   string(res.UID),
		Memory: res.Status.Allocatable.Memory().String(),
		CPU:    res.Status.Allocatable.Cpu().String(),
	}, err

}

func (s nodeService) GetSingleNodeMetrics(name string) (Metrics, error) {
	res, err := s.client.GetNodeMetrics(name)
	if err != nil {
		return Metrics{}, err
	}

	return Metrics{
		Node: Node{
			Name:   res.Name,
			UUID:   string(res.UID),
			Memory: res.Usage.Memory().String(),
			CPU:    res.Usage.Cpu().String(),
		},
	}, err

}

func (s nodeService) GetNodesMetrics() ([]Metrics, error) {
	var metrics []Metrics
	res, err := s.client.GetNodesMetrics()
	if err != nil {
		return metrics, err
	}

	for _, i := range res.Items {

		metrics = append(metrics,
			Metrics{
				Node: Node{
					Name:   i.Name,
					UUID:   string(i.UID),
					Memory: i.Usage.Memory().String(),
					CPU:    i.Usage.Cpu().String(),
				},
			})
	}

	return metrics, err
}
