package autoscaler

import (
	"errors"

	"github.com/gorilla/mux"
	v2 "k8s.io/api/autoscaling/v2"
	coreV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	res, err := s.client.GetAutoscaler(name)
	if err != nil {
		return PodAutoscaler{}, err
	}

	return convertToPodAutoscaler(res), err
}

func (s autoscalerService) GetAutoscalers() (autoscalers []PodAutoscaler, err error) {
	res, err := s.client.GetAutoscalers()
	if err != nil {
		return
	}

	for _, i := range res.Items {
		autoscalers = append(autoscalers, convertToPodAutoscaler(&i))
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

func (s autoscalerService) CreateAutoscaler(createAutoscaler PodAutoscalerCreateRequest) (PodAutoscaler, error) {
	var podAutoscaler PodAutoscaler
	req, err := prepareToCreate(createAutoscaler)
	if err != nil {
		return podAutoscaler, err
	}

	k8sAutoScaler, err := s.client.CreateAutoscaler(req)
	if err != nil {
		return podAutoscaler, err
	}

	return convertToPodAutoscaler(k8sAutoScaler), err
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
				kr.Resource.Target.AverageUtilization = &r.Value
			}
		}
	}
}

func convertToPodAutoscaler(k8sAutoscaler *v2.HorizontalPodAutoscaler) PodAutoscaler {
	var podResources []PodResources

	for _, r := range k8sAutoscaler.Spec.Metrics {
		podResources = append(podResources, PodResources{
			Name:  string(r.Resource.Name),
			Type:  string(r.Resource.Target.Type),
			Value: r.Resource.Target.Value.String(),
		})
	}

	return PodAutoscaler{
		Name:      k8sAutoscaler.Name,
		ID:        string(k8sAutoscaler.UID),
		Resources: podResources,
		Replicas: PodReplicas{
			Min: *k8sAutoscaler.Spec.MinReplicas,
			Max: k8sAutoscaler.Spec.MaxReplicas,
		},
	}
}

func prepareToCreate(createAutoscaler PodAutoscalerCreateRequest) (*v2.HorizontalPodAutoscaler, error) {
	metrics := make([]v2.MetricSpec, len(createAutoscaler.Resources))
	for _, r := range createAutoscaler.Resources {
		if r.Name != "cpu" || r.Name == "memory" {
			return &v2.HorizontalPodAutoscaler{}, errors.New("only accepted cpu or memory resource")
		}

		metrics = append(metrics, v2.MetricSpec{
			Type: v2.ResourceMetricSourceType,
			Resource: &v2.ResourceMetricSource{
				Name: coreV1.ResourceName(createAutoscaler.Name),
				Target: v2.MetricTarget{
					Type:               v2.UtilizationMetricType,
					AverageUtilization: &r.Value,
				},
			},
		})
	}

	return &v2.HorizontalPodAutoscaler{
		ObjectMeta: v1.ObjectMeta{
			Name: createAutoscaler.Name,
		},
		Spec: v2.HorizontalPodAutoscalerSpec{
			ScaleTargetRef: v2.CrossVersionObjectReference{
				Kind: "Deployment",
				Name: createAutoscaler.DeploymentTarget,
			},
			Metrics: metrics,
		},
	}, nil
}
