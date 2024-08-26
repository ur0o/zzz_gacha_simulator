package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zzz_gacha/models"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, models.Draw())
}
