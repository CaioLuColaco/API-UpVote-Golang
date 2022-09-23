package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowCoins(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id": 1,
		"coin": "Bitcoin",
		"votos": 3,
	})
}

