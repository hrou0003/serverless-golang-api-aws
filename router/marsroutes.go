package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hrou0003/serverless-golang-api-aws/mars"
)

func marsRoutes(r *gin.Engine) {
	marsRoutes := r.Group("/mars")
	{
		marsRoutes.GET("/weather/sols", func(ctx *gin.Context) {
			mars.GetMarsSolsWeatherInfoController(ctx)
		})
	}
}
