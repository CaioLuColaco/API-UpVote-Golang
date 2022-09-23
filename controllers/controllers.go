package controllers

import (
	"net/http"

	"github.com/CaioLuColaco/api-upVote-golang/database"
	"github.com/CaioLuColaco/api-upVote-golang/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"github.com/gin-gonic/gin"
)


// ShowAllCoins godoc
// @Summary      Show all coins registred
// @Description  Route used to list all coins registred
// @Tags         coin
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Coin
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /coins [get]
func ShowAllCoins(c *gin.Context) {
	var coins []models.Coin
	database.DB.Find(&coins)
	c.JSON(http.StatusOK, coins)
}

// ShowOneCoin godoc
// @Summary      Show one coin registred 
// @Description  Route used to get a one coin registred by ID
// @Tags         coin
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Coin ID"
// @Success      200  {object}  models.Coin
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /coin/{id} [get]
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

// ShowCoinByName godoc
// @Summary      Show one coin registred 
// @Description  Route used to get a one coin registred by Name
// @Tags         coin
// @Accept       json
// @Produce      json
// @Param        name   path      string  true  "Coin Name"
// @Success      200  {object}  models.Coin
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /coin/name/{name} [get]
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

// ShowCoinByShortName godoc
// @Summary      Show one coin registred 
// @Description  Route used to get a one coin registred by ShortName
// @Tags         coin
// @Accept       json
// @Produce      json
// @Param        shortname   path      string  true  "Coin Name"
// @Success      200  {object}  models.Coin
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /coins/shortName/{shortname} [get]
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

// CreateCoin godoc
// @Summary      Create one coin
// @Description  Route used to create a one coin
// @Tags         coin
// @Accept       json
// @Produce      json
// @Param        coin   body    models.Coin  true  "Model of Coin"
// @Success      200  {object}  models.Coin
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /coin [post]
func CreateCoin(c *gin.Context) {
	var coin models.Coin

	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.ValidateCoinData(&coin); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&coin)

	c.JSON(http.StatusOK, coin)
}


// UpdateCoin godoc
// @Summary      Create one coin
// @Description  Route used to update a one coin by ID
// @Tags         coin
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Coin ID"
// @Param        coin   body    models.Coin  true  "Model of Coin"
// @Success      200  {object}  models.Coin
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /coin/{id} [patch]
func UpdateCoin(c *gin.Context) {
	var coin models.Coin
	id := c.Params.ByName("id")
	database.DB.First(&coin, id)

	if coin.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Coin not found"})
		return
	}
	
	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
		
	if err := models.ValidateCoinData(&coin); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	


	database.DB.Save(&coin)
	c.JSON(http.StatusOK, coin)
}

// DeleteCoin godoc
// @Summary      Delete one coin
// @Description  Route used to delete a one coin by ID
// @Tags         coin
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Coin ID"
// @Success      200  {object}  models.Coin
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Router       /coin/{id} [delete]
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