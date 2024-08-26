package config

import (
	"github.com/gin-gonic/gin"

	"zzz_gacha/controllers"
)

func SetRouting(e *gin.Engine) {
	indices := e.Group("/api")
	{
		indices.GET("/", controllers.Index)
	}
}
