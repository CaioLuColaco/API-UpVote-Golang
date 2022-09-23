package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/api-upVote-golang/database"
	"github.com/CaioLuColaco/api-upVote-golang/models"
	"github.com/gin-gonic/gin"
)

func ShowAllCoins(c *gin.Context) {
	var coins []models.Coin
	database.DB.Find(&coins)
	c.JSON(http.StatusOK, coins)
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