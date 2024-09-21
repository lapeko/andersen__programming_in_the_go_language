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
	article := &models.Article{}
	if err := c.ShouldBind(article); err != nil {
		sendError(c, http.StatusBadRequest)
	}
	if err := a.repo.Create(article); err != nil {
		sendError(c, http.StatusInternalServerError)
	}
	sendSuccessWithStatus(c, http.StatusCreated, article)
}

func (a *articlesHandler) GetAll(c *gin.Context) {
	var articles []models.Article
	if err := a.repo.GetAll(&articles); err != nil {
		sendError(c, http.StatusInternalServerError)
	}
	sendSuccess(c, &articles)
}

func (a *articlesHandler) GetById(c *gin.Context) {
}

func (a *articlesHandler) Update(c *gin.Context) {
}

func (a *articlesHandler) Delete(c *gin.Context) {
}
