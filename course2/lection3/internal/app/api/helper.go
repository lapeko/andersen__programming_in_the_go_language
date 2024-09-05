package api

import (
	"encoding/json"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/storage"
	"net/http"
)

const apiPrefix = "/api/v1"

type response struct {
	Ok    bool        `json:"ok"`
	Body  interface{} `json:"body"`
	Error string      `json:"error"`
}

func (a *API) configureRouting() {
	a.router.HandleFunc(apiPrefix+"/ping", func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = writer.Write([]byte("Pong"))
	})
	a.router.HandleFunc(apiPrefix+"/users", a.GetAllUsers)
	a.router.HandleFunc(apiPrefix+"/articles", a.GetAllArticles)
}

func (a *API) configureStorage() error {
	a.storage = storage.New(a.config.StorageConfig)
	return a.storage.Open()
}

func sendError(w http.ResponseWriter, err error, statusCode int) {
	log := logger.Get()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response{false, nil, err.Error()}); err != nil {
		log.Error(err)
	}
}

func sendSuccess(w http.ResponseWriter, body interface{}) {
	sendSuccessWithCode(w, body, http.StatusOK)
}

func sendSuccessWithCode(w http.ResponseWriter, body interface{}, statusCode int) {
	log := logger.Get()
	w.Header().Set("Content-Type", "application/json")
	if statusCode != 0 {
		w.WriteHeader(statusCode)
	}
	if err := json.NewEncoder(w).Encode(response{true, body, ""}); err != nil {
		log.Error(err)
	}
}
