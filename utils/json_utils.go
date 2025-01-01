package utils

import (
	"encoding/json"
	"net/http"
)

func DecodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}

func WriteJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
