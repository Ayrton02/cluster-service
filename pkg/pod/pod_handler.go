package pod

import (
	"net/http"
	"service-client/middleware"

	"github.com/gorilla/mux"
)

func GetSinglePod(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]
	pod, err := Service.GetSinglePod(name)

	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, pod)
}

func GetPods(w http.ResponseWriter, r *http.Request) {
	pods, err := Service.GetPods("")
	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusOK, pods)
}

func DeleteSinglePod(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	name := pathParams["name"]
	err := Service.DeleteSinglePod(name)

	if err != nil {
		middleware.WriteResponse(w, http.StatusInternalServerError, err)
	}

	middleware.WriteResponse(w, http.StatusNoContent, err)
}
