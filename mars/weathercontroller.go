package mars

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMarsSolsWeatherInfoController(ctx *gin.Context) {
	response, err := GetAllSolsWeather()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, response)

}
