package api

import (
	"encoding/json"
	"errors"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/internal/app/models"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
	"net/http"
)

func (a *API) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()

	users, err := a.storage.Users().GetAll()

	if err != nil {
		log.Error(err)
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	log.Debugln("received users:", users)

	sendSuccess(w, users)
}

func (a *API) CreateUser(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		log.Error(err)
		sendError(w, err, http.StatusBadRequest)
		return
	}

	user, err = a.storage.Users().CreateUser(user)

	if err != nil {
		log.Error(err)
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	if user == nil {
		log.Error("email already taken")
		sendError(w, errors.New("email already taken"), http.StatusConflict)
		return
	}

	log.Debugln("received user:", user)

	sendSuccess(w, user)
}

func (a *API) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()

	articles, err := a.storage.Articles().GetAll()

	if err != nil {
		log.Error(err)
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	log.Debugln("received articles:", articles)

	sendSuccess(w, articles)
}
