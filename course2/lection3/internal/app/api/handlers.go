package api

import (
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
