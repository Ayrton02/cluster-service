package node

import (
	"net/http"
	"service-client/middleware"

	"github.com/gorilla/mux"
)

func GetSingleNode(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	deployment, err := Service.GetSingleNode(name)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, deployment)
}

func GetNodes(w http.ResponseWriter, r *http.Request) {

	deployments, err := Service.GetNodes()
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, deployments)
}

func GetSingleNodeMetrics(w http.ResponseWriter, r *http.Request) {

	deployments, err := Service.GetNodes()
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, deployments)
}

func GetNodesMetrics(w http.ResponseWriter, r *http.Request) {

	deployments, err := Service.GetNodes()
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, deployments)
}
