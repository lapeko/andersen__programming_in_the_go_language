package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/internal/handlers"
	"github.com/lapeko/andersen__programming_in_the_go_language/course2/lection6/example2-gorm-gin/repository"
)

func SetupArticlesRouter(path string, router *gin.Engine, repo *repository.ArticlesRepo) {
	r := router.Group(path)
	handler := handlers.GetHandlers(repo)
	{
		r.POST("", handler.Articles.Create)
		r.GET("", handler.Articles.GetAll)
		r.GET("/:id", handler.Articles.GetById)
		r.PUT("", handler.Articles.Update)
		r.DELETE("/:id", handler.Articles.Delete)
	}
}
