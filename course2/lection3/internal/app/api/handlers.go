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
	log.Debug("GetAllUsers")

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
	log.Debug("CreateUser")

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
	log.Debug("GetAllArticles")

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
	log.Debug("GetArticle")

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		log.Error("no id in params")
		sendError(w, errors.New("no ID found"), http.StatusInternalServerError)
		return
	}

	digitId, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		log.Error("incorrect ID format")
		sendError(w, errors.New("incorrect ID format"), http.StatusBadRequest)
		return
	}

	article, err := a.storage.Articles().GetById(int(digitId))

	if err != nil {
		log.Error(err)
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	if article == nil {
		log.Debugln("article not found")
		sendError(w, errors.New("article not found"), http.StatusNotFound)
		return
	}

	sendSuccess(w, article)
}

func (a *API) CreateArticle(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	log.Debug("CreateArticle")

	article := &models.Article{}
	err := json.NewDecoder(r.Body).Decode(article)

	if err != nil {
		log.Error("request body is invalid")
		sendError(w, err, http.StatusBadRequest)
		return
	}

	article, err = a.storage.Articles().Create(article)

	if err != nil {
		log.Error("article creation error: ", err)
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	sendSuccessWithCode(w, article, http.StatusCreated)
}

func (a *API) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	log.Debug("DeleteArticle")

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		log.Error("no id in params")
		sendError(w, errors.New("no ID found"), http.StatusInternalServerError)
		return
	}

	digitId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		log.Error("incorrect ID format")
		sendError(w, errors.New("incorrect ID format"), http.StatusBadRequest)
		return
	}

	ok, err = a.storage.Articles().Delete(digitId)

	if err != nil {
		log.Error(err)
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	if !ok {
		log.Debugln("article does not exist")
		sendError(w, errors.New("article not found"), http.StatusNotFound)
		return
	}

	sendSuccessWithCode(w, digitId, http.StatusOK)
}

func (a *API) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	log := logger.Get()
	log.Debug("UpdateArticle")

	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		log.Error("no id in params")
		sendError(w, errors.New("no ID found"), http.StatusInternalServerError)
		return
	}

	digitId, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		log.Error("incorrect ID format")
		sendError(w, errors.New("incorrect ID format"), http.StatusBadRequest)
		return
	}

	article := &models.Article{Id: uint(digitId)}
	err = json.NewDecoder(r.Body).Decode(article)

	if err != nil {
		log.Error("request body is invalid")
		sendError(w, err, http.StatusBadRequest)
		return
	}

	article, err = a.storage.Articles().Update(article)

	if err != nil {
		log.Error("update article db error: ", err)
		sendError(w, err, http.StatusInternalServerError)
		return
	}

	if article == nil {
		log.Error("article not found")
		sendError(w, errors.New("not found"), http.StatusNotFound)
		return
	}

	sendSuccessWithCode(w, article, http.StatusOK)
}
