package autoscaler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func registerAutoscalerRoutes(router *mux.Router) {
	router.HandleFunc("/autoscalers/{name}", GetSingleAutoscaler).Methods(http.MethodGet)
	router.HandleFunc("/autoscalers", GetAutoscalers).Methods(http.MethodGet)
	router.HandleFunc("/autoscalers/{name}", UpdateAutoscaler).Methods(http.MethodPut)
	router.HandleFunc("/autoscalers", CreateAutoscaler).Methods(http.MethodPost)
}
