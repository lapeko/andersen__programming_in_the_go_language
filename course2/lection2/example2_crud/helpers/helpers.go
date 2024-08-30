package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Ok           bool
	Body         interface{}
	ErrorMessage string
}

func SendJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response{Ok: true, Body: payload, ErrorMessage: ""})

	if err != nil {
		SendError(w, "Serialization error: "+err.Error(), http.StatusInternalServerError)
	}
}

func SendError(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response{
		Ok:           false,
		Body:         nil,
		ErrorMessage: errorMessage,
	})
}
