package node

import (
	"net/http"
	"service-client/middleware"

	"github.com/gorilla/mux"
)

func GetSingleNode(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	node, err := Service.GetSingleNode(name)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, node)
}

func GetNodes(w http.ResponseWriter, r *http.Request) {

	nodes, err := Service.GetNodes()
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, nodes)
}

func GetSingleNodeMetrics(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	node, err := Service.GetSingleNodeMetrics(name)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, node)
}

func GetNodesMetrics(w http.ResponseWriter, r *http.Request) {

	nodes, err := Service.GetNodesMetrics()
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, nodes)
}
