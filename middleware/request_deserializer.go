package middleware

import (
	"encoding/json"
	"io"
)

func DeserializeJson(body io.ReadCloser, object interface{}) error {
	return json.NewDecoder(body).Decode(object)
}
