package pod

import (
	"net/http"

	"github.com/gorilla/mux"
)

func registerPodRoutes(router *mux.Router) {
	router.HandleFunc("/pods/{name}", GetSinglePod).Methods(http.MethodGet)
	router.HandleFunc("/pods", GetPods).Methods(http.MethodGet)
	router.HandleFunc("/pods/{name}", DeleteSinglePod).Methods(http.MethodDelete)
}
