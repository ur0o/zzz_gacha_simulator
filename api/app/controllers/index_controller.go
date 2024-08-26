package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin_api/app/models"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, models.Draw())
}
