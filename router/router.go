package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRouter(r *gin.Engine) {
	marsRoutes(r)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
