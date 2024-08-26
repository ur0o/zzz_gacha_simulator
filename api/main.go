package main

import (
	"github.com/gin-gonic/gin"
	"zzz_gacha/config"
	"zzz_gacha/database"
	"zzz_gacha/models"
)

func main() {
	engine := gin.Default()

	database.InitDB()
	models.Migrate()

	config.SetHeader(engine)
	config.SetRouting(engine)
	engine.Run(":8080")
}
