package middleware

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, status int, data interface{}) {
	res, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Write(res)
}
