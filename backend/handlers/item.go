package handlers

import (
	"backend/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ListItems(c *gin.Context) {
	var items []models.Item
	models.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}