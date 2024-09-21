package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/internal/models"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/repository"
	"net/http"
)

type Handlers struct {
	Articles *articlesHandler
}

type articlesHandler struct {
	repo *repository.ArticlesRepo
}

var hs *Handlers

func GetHandlers(r *repository.ArticlesRepo) *Handlers {
	if hs != nil {
		return hs
	}
	hs = &Handlers{Articles: &articlesHandler{repo: r}}
	return hs
}

func (a *articlesHandler) Create(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBind(&article); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
	}
	if err := a.repo.Create(&article); err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccessWithStatus(c, http.StatusCreated, &article)
}

func (a *articlesHandler) GetAll(c *gin.Context) {
	var articles []models.Article
	if err := a.repo.GetAll(&articles); err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	sendSuccess(c, &articles)
}

func (a *articlesHandler) GetById(c *gin.Context) {

	var article models.Article
	if err := a.repo.GetById(&article, c.Param("id")); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(c, &article)
}

func (a *articlesHandler) Update(c *gin.Context) {
	var article models.Article
	if err := c.ShouldBind(&article); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.repo.Update(&article); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	sendSuccess(c, &article)
}

func (a *articlesHandler) Delete(c *gin.Context) {
	if err := a.repo.DeleteById(c.Param("id")); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
	}

	sendSuccess(c, nil)
}
