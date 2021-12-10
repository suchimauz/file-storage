package http

import (
	"github.com/gin-gonic/gin"
)

type dataResponse struct {
	Url string `json:"url"`
}

type idResponse struct {
	ID interface{} `json:"id"`
}

type response struct {
	Message string `json:"message"`
}

func newResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, response{message})
}
