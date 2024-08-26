package main

import (
	"github.com/gin-gonic/gin"
	"gin_api/config"
)

func main() {
	engine := gin.Default()

	config.SetHeader(engine)
	config.SetRouting(engine)
	config.InitDB()
	engine.Run(":8080")
}
