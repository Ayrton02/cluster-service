package middleware

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, data interface{}) {
	res, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	w.Write(res)
}
