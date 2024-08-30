package helpers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func ParseParamId(r *http.Request) (uint, error) {
	vars := mux.Vars(r)

	idString, ok := vars["id"]

	if !ok {
		return 0, errors.New("ID does not exist")
	}

	idUint64, err := strconv.ParseUint(idString, 10, 0)

	if err != nil {
		return 0, errors.New("ID is not valid")
	}

	return uint(idUint64), nil
}
