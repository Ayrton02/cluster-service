package main

import (
	"context"
	"os"
	"service-client/configs"
	"service-client/pkg/app"
)

func main() {
	port := os.Getenv("PORT")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	namespace := os.Getenv("NAMESPACE")
	metricsNamespace := os.Getenv("METRICS_NAMESPACE")

	api := configs.NewAPIConfig(port, allowedOrigins)
	k8s, err := configs.NewKubernetesConfig(namespace, metricsNamespace)

	if err != nil {
		panic(err)
	}

	config := configs.Config{
		API: *api,
		K8S: *k8s,
	}

	app := app.App{
		Ctx:    context.Background(),
		Config: config,
	}

	app.Start()
}
