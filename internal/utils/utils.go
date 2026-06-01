package utils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, code int, params interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(params)
}
