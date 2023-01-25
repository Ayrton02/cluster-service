package deployment

import (
	"net/http"

	"github.com/gorilla/mux"
)

func registerDeploymentRoutes(router *mux.Router) {
	router.HandleFunc("/deployments/{name}", GetSingleDeployment).Methods(http.MethodGet)
	router.HandleFunc("/deployments", GetDeployments).Methods(http.MethodGet)
	router.HandleFunc("/deployments/{name}", DeleteSingleDeployment).Methods(http.MethodDelete)
	router.HandleFunc("/deployments/{name}", UpdateSingleDeployment).Methods(http.MethodPut)
}
