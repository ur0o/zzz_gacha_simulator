package config

import (
	"github.com/gin-gonic/gin"

	"zzz_gacha/controllers"
)

func SetRouting(e *gin.Engine) {
	indices := e.Group("/api")
	gachas := indices.Group("/gacha")
	{
		gachas.GET("/", controllers.IndexGacha)
		gachas.GET("/:type/:id", controllers.GetGacha)
		gachas.GET("/:type/:id/draw", controllers.DrawGacha)
	}
}
