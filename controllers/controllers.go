package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/api-upVote-golang/database"
	"github.com/CaioLuColaco/api-upVote-golang/models"
	"github.com/gin-gonic/gin"
)

func ShowCoins(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id": 1,
		"coin": "Bitcoin",
		"votos": 3,
	})
}

func CreateCoin(c *gin.Context) {
	var coin models.Coin
	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&coin)
	c.JSON(http.StatusOK, coin)
}