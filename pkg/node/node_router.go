package node

import (
	"net/http"

	"github.com/gorilla/mux"
)

func registerNodeRoutes(router *mux.Router) {
	router.HandleFunc("/nodes/{name}", GetSingleNode).Methods(http.MethodGet)
	router.HandleFunc("/nodes", GetNodes).Methods(http.MethodGet)
	router.HandleFunc("/nodes/metrics", GetNodesMetrics).Methods(http.MethodGet)
	router.HandleFunc("/nodes/{name}/metrics", GetSingleNodeMetrics).Methods(http.MethodGet)
}
