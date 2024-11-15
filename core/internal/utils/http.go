package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func DecodeJson(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
