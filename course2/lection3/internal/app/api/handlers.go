package api

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/internal/app/models"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection3/pkg/logger"
	"net/http"
	"strconv"
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

	sendSuccessWithCode(w, user, http.StatusCreated)
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

func (a *API) GetArticle(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		log.Error("No id in params")
		sendError(w, errors.New("no ID found"), http.StatusInternalServerError)
		return
	}

	digitId, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		log.Error("incorrect ID format")
		sendError(w, errors.New("incorrect ID format"), http.StatusBadRequest)
		return
	}

	article, err := a.storage.Articles().GetArticleById(int(digitId))

	if err != nil {
		log.Error(err)
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	if article == nil {
		log.Debugln("user not found")
		sendError(w, errors.New("article not found"), http.StatusNotFound)
		return
	}

	sendSuccess(w, article)
}

// Create, Delete, Update
