package autoscaler

import (
	"net/http"
	"service-client/middleware"

	"github.com/gorilla/mux"
)

func GetSingleAutoscaler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	autoscaler, err := Service.GetSingleAutoscaler(name)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, autoscaler)
}

func GetAutoscalers(w http.ResponseWriter, r *http.Request) {
	autoscaler, err := Service.GetAutoscalers()
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, autoscaler)
}

func UpdateAutoscaler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]

	var updateAutoscalerRequest PodAutoscalerUpdateRequest
	err := middleware.DeserializeJson(r.Body, &updateAutoscalerRequest)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	updateAutoscalerRequest.Name = name

	err = Service.UpdateAutoscaler(updateAutoscalerRequest)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusNoContent, err)

}

func CreateAutoscaler(w http.ResponseWriter, r *http.Request) {
	var createAutoscalerRequest PodAutoscalerCreateRequest
	err := middleware.DeserializeJson(r.Body, &createAutoscalerRequest)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	res, err := Service.CreateAutoscaler(createAutoscalerRequest)
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusCreated, res)

}
