package app

import (
	"context"
	"net/http"
	"service-client/configs"
	"service-client/pkg/cluster"
	"service-client/pkg/deployment"
	"service-client/pkg/node"
	"service-client/pkg/pod"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Ctx    context.Context
	Config configs.Config
}

func (a *App) Start() {
	namespace := a.Config.K8S.Namespace
	metricsNamespace := a.Config.K8S.MetricsNamespace

	k8sclient := cluster.NewK8SClient(
		a.Ctx,
		cluster.K8SClientSets{
			Client:  a.Config.K8S.ClientSet,
			Metrics: *a.Config.K8S.Metrics,
		},
		cluster.K8SNamespaces{
			Operational: namespace,
			Metrics:     metricsNamespace,
		},
	)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{a.Config.API.AllowedOrigins})

	router := mux.NewRouter()
	pod.InitPodService(k8sclient.Pods, router)
	deployment.InitDeploymentService(k8sclient.Deployment, router)
	node.InitNodeService(k8sclient.Node, router)
	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
