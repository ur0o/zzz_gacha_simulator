package main

import (
	"github.com/gin-gonic/gin"
	"zzz_gacha/config"
	"zzz_gacha/database"
)

func main() {
	engine := gin.Default()

	database.InitDB()

	config.SetHeader(engine)
	config.SetRouting(engine)
	engine.Run(":8080")
}
