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

func ShowOneCoin(c *gin.Context) {
	var coin models.Coin
	id := c.Params.ByName("id")
	database.DB.First(&coin, id)
	
	if coin.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Coin not found",
		})
		return
	}

	c.JSON(http.StatusOK, coin)
}

func ShowCoinByName(c *gin.Context) {
	var coin models.Coin
	name := c.Params.ByName("name")
	database.DB.Where(&models.Coin{Name: name}).First(&coin)
	
	if coin.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Coin not found",
		})
		return
	}

	c.JSON(http.StatusOK, coin)
}

func ShowCoinByShortName(c *gin.Context) {
	var coin models.Coin
	shortname := c.Params.ByName("shortname")
	database.DB.Where(&models.Coin{ShortName: shortname}).First(&coin)
	
	if coin.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Coin not found",
		})
		return
	}

	c.JSON(http.StatusOK, coin)
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

func UpdateCoin(c *gin.Context) {
	var coin models.Coin

	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	id := c.Params.ByName("id")
	database.DB.First(&coin, id)
	if coin.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Coin not found",
		})
		return
	}

	database.DB.Model(&coin).UpdateColumns(coin)
	c.JSON(http.StatusOK, coin)
}

func DeleteCoin(c *gin.Context) {
	var coin models.Coin
	id := c.Params.ByName("id")
	database.DB.First(&coin, id)

	if coin.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Coin not found",
		})
		return
	}

	database.DB.Delete(&coin, id)
	c.JSON(http.StatusOK, gin.H{"data": "Coin deleted!"})
}