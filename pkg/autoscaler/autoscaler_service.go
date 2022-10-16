package autoscaler

import (
	"github.com/gorilla/mux"
	v2 "k8s.io/api/autoscaling/v2"
	"k8s.io/apimachinery/pkg/api/resource"
)

var Service *autoscalerService

type autoscalerService struct {
	client autoscalerClient
}

func InitAutoscalertService(client autoscalerClient, route *mux.Router) {
	registerAutoscalerRoutes(route)
	Service = &autoscalerService{
		client: client,
	}
}

func (s autoscalerService) GetSingleAutoscaler(name string) (PodAutoscaler, error) {
	var podAutoscaler PodAutoscaler
	var podResources []PodResources

	res, err := s.client.GetAutoscaler(name)
	if err != nil {
		return podAutoscaler, err
	}

	for _, r := range res.Spec.Metrics {
		podResources = append(podResources, PodResources{
			Name:  string(r.Resource.Name),
			Type:  string(r.Resource.Target.Type),
			Value: r.Resource.Target.Value.String(),
		})
	}

	podAutoscaler = PodAutoscaler{
		Name: res.Name,
		ID:   string(res.UID),
		Replicas: PodReplicas{
			Min: *res.Spec.MinReplicas,
			Max: res.Spec.MaxReplicas,
		},
		Resources: podResources,
	}

	return podAutoscaler, err
}

func (s autoscalerService) GetAutoscalers() (autoscalers []PodAutoscaler, err error) {
	res, err := s.client.GetAutoscalers()
	if err != nil {
		return
	}

	for _, i := range res.Items {
		resources := make([]PodResources, 0)
		for _, r := range i.Spec.Metrics {
			resources = append(resources, PodResources{
				Name:  string(r.Resource.Name),
				Type:  string(r.Resource.Target.Type),
				Value: r.Resource.Target.Value.String(),
			})
		}

		autoscalers = append(autoscalers, PodAutoscaler{
			Name: i.Name,
			ID:   string(i.UID),
			Replicas: PodReplicas{
				Min: *i.Spec.MinReplicas,
				Max: i.Spec.MaxReplicas,
			},
			Resources: resources,
		})
	}

	return
}

func (s autoscalerService) UpdateAutoscaler(updateAutoscaler PodAutoscalerUpdateRequest) error {
	k8sAutoscaler, err := s.client.GetAutoscaler(updateAutoscaler.Name)
	if err != nil {
		return err
	}

	prepareToUpdate(k8sAutoscaler, updateAutoscaler)
	_, err = s.client.UpdateAutoscaler(k8sAutoscaler)

	return err
}

func prepareToUpdate(k8sAutoscaler *v2.HorizontalPodAutoscaler, updateAutoscaler PodAutoscalerUpdateRequest) {
	if updateAutoscaler.Replicas.Max > 0 {
		k8sAutoscaler.Spec.MaxReplicas = updateAutoscaler.Replicas.Max
	}

	if updateAutoscaler.Replicas.Min > 0 {
		k8sAutoscaler.Spec.MinReplicas = &updateAutoscaler.Replicas.Min
	}

	for _, r := range updateAutoscaler.Resources {
		for _, kr := range k8sAutoscaler.Spec.Metrics {
			if kr.Resource.Name.String() == r.Name {
				kr.Resource.Target.Value = &resource.Quantity{Format: resource.Format(r.Value)}
			}
		}
	}
}
