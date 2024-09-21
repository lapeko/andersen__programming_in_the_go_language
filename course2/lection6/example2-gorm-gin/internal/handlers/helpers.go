package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Ok         bool        `json:"ok"`
	Body       interface{} `json:"body"`
}

func sendSuccess(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, Response{Ok: true, StatusCode: http.StatusOK, Body: body})
}

func sendSuccessWithStatus(c *gin.Context, statusCode int, body interface{}) {
	c.JSON(statusCode, Response{Ok: true, StatusCode: statusCode, Body: body})
}

func sendError(c *gin.Context, statusCode int) {
	c.JSON(statusCode, Response{Ok: false, StatusCode: statusCode, Body: nil})
}
