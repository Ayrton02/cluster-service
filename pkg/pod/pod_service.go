package pod

import (
	"github.com/gorilla/mux"
)

var Service *podService

type podService struct {
	client podClient
}

func InitPodService(client podClient, route *mux.Router) {
	registerPodRoutes(route)
	Service = &podService{
		client: client,
	}
}

func (s podService) GetPods(options string) ([]Pod, error) {
	var pods []Pod

	res, err := s.client.GetPods(options)
	if err != nil {
		return pods, err
	}

	for _, p := range res.Items {
		var metrics Metrics
		podMetrics, err := s.client.GetPodMetrics(p.Name)
		if err == nil {
			metrics = Metrics{
				Memory: podMetrics.Containers[0].Usage.Memory().String(),
				CPU:    podMetrics.Containers[0].Usage.Cpu().String(),
			}
		}

		pods = append(pods,
			Pod{
				Name:           p.Name,
				Status:         string(p.Status.Phase),
				IP:             p.Status.PodIP,
				ContainerImage: p.Spec.Containers[0].Image,
				ContainerPort:  p.Spec.Containers[0].Ports[0].HostPort,
				UUID:           string(p.UID),
				Metrics:        metrics,
			},
		)
	}

	return pods, nil
}

func (s podService) GetSinglePod(name string) (Pod, error) {
	var pod Pod

	res, err := s.client.GetPod(name)

	if err != nil {
		return pod, err
	}

	podMetrics, err := s.client.GetPodMetrics(res.Name)
	if err != nil {
		return pod, err
	}

	pod = Pod{
		Name:           res.Name,
		Status:         string(res.Status.Phase),
		IP:             res.Status.PodIP,
		UUID:           string(res.UID),
		ContainerImage: res.Spec.Containers[0].Image,
		ContainerPort:  res.Spec.Containers[0].Ports[0].HostPort,
		Metrics: Metrics{
			Memory: podMetrics.Containers[0].Usage.Memory().String(),
			CPU:    podMetrics.Containers[0].Usage.Cpu().String(),
		},
	}

	return pod, err
}

func (s podService) DeleteSinglePod(name string) error {
	return s.client.DeletePod(name)
}
