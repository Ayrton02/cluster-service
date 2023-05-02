package deployment

import (
	"net/http"
	"service-client/middleware"

	"github.com/gorilla/mux"
)

func GetSingleDeployment(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	deployment, err := Service.GetSingleDeployment(name)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, deployment)
}

func GetDeployments(w http.ResponseWriter, r *http.Request) {

	deployments, err := Service.GetDeployments()
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, deployments)
}

func DeleteSingleDeployment(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	err := Service.DeleteSingleDeployment(name)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusNoContent, err)
}

func UpdateSingleDeployment(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	var updateDeploymentRequest DeploymentUpdateRequest
	err := middleware.DeserializeJson(r.Body, &updateDeploymentRequest)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	updateDeploymentRequest.Name = name

	err = Service.UpdateSingleDeployment(updateDeploymentRequest)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusNoContent, err)
}

func CreateDeployment(w http.ResponseWriter, r *http.Request) {
	var createDeployment DeploymentCreateRequest
	err := middleware.DeserializeJson(r.Body, &createDeployment)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	deployment, err := Service.CreateDeployment(createDeployment)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusCreated, deployment)
}
