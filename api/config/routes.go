package config

import (
	"github.com/gin-gonic/gin"

	"zzz_gacha/controllers"
)

func SetRouting(e *gin.Engine) {
	indices := e.Group("/api")
	gachas := indices.Group("/gacha")
	{
		gachas.GET("/:id", controllers.GetGacha)
		gachas.GET("/:id/draw", controllers.DrawGacha)
	}
}
