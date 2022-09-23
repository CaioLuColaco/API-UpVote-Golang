package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/CaioLuColaco/api-upVote-golang/controllers"
	"github.com/CaioLuColaco/api-upVote-golang/database"
	"github.com/CaioLuColaco/api-upVote-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int
var NAME string
var SHORTNAME string

func CreateCoin() {
	coin := models.Coin{Name: "CoinX", ShortName: "CNX", Votes: 1, Price: 2.50}
	database.DB.Create(&coin)
	ID = int(coin.ID)
	NAME = coin.Name
	SHORTNAME = coin.ShortName
}

func DeleteCoin() {
	var coin models.Coin
	database.DB.Delete(&coin, ID)
}

func SetUpRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestShowAllCoinsHandler(t *testing.T) {
	database.ConnectToDatabase()
	CreateCoin()
	defer DeleteCoin()
	r := SetUpRoutesTest()
	r.GET("/coins", controllers.ShowAllCoins)
	req, _ := http.NewRequest("GET", "/coins", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")
}

func TestShowOneCoinHandler(t *testing.T) {

	database.ConnectToDatabase()
	
	CreateCoin()
	defer DeleteCoin()
	
	r := SetUpRoutesTest()
	r.GET("/coins/:id", controllers.ShowOneCoin)
	routeTest := `/coins/` + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", routeTest, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var coinTest models.Coin
	json.Unmarshal(res.Body.Bytes(), &coinTest)

	assert.Equal(t, "CoinX", coinTest.Name, "The values must be the same")
	assert.Equal(t, "CNX", coinTest.ShortName, "The values must be the same")
	assert.Equal(t, 1, coinTest.Votes, "The values must be the same")
	assert.Equal(t, 2.50, coinTest.Price, "The values must be the same")
	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")
}

func TestShowCoinByNameHandler(t *testing.T) {

	database.ConnectToDatabase()
	
	CreateCoin()
	defer DeleteCoin()
	
	r := SetUpRoutesTest()
	r.GET("/coins/name/:name", controllers.ShowCoinByName)
	routeTest := `/coins/name/` + NAME
	req, _ := http.NewRequest("GET", routeTest, nil)
	res := httptest.NewRecorder()
	
	r.ServeHTTP(res, req)
	
	var coinTest models.Coin
	json.Unmarshal(res.Body.Bytes(), &coinTest)

	assert.Equal(t, "CoinX", coinTest.Name, "The values must be the same")
	assert.Equal(t, "CNX", coinTest.ShortName, "The values must be the same")
	assert.Equal(t, 1, coinTest.Votes, "The values must be the same")
	assert.Equal(t, 2.50, coinTest.Price, "The values must be the same")
	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")
}

func TestShowCoinByShortNameHandler(t *testing.T) {

	database.ConnectToDatabase()
	
	CreateCoin()
	defer DeleteCoin()
	
	r := SetUpRoutesTest()
	r.GET("/coins/shortName/:shortname", controllers.ShowCoinByShortName)
	routeTest := `/coins/shortName/` + SHORTNAME
	req, _ := http.NewRequest("GET", routeTest, nil)
	res := httptest.NewRecorder()
	
	r.ServeHTTP(res, req)
	
	var coinTest models.Coin
	json.Unmarshal(res.Body.Bytes(), &coinTest)

	assert.Equal(t, "CoinX", coinTest.Name, "The values must be the same")
	assert.Equal(t, "CNX", coinTest.ShortName, "The values must be the same")
	assert.Equal(t, 1, coinTest.Votes, "The values must be the same")
	assert.Equal(t, 2.50, coinTest.Price, "The values must be the same")
	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")
}

func TestDeleteCoinHandler(t * testing.T) {

	database.ConnectToDatabase()
	
	CreateCoin()

	r := SetUpRoutesTest()
	r.DELETE("/coin/:id", controllers.DeleteCoin)
	routeTest := `/coin/` + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", routeTest, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")

}

func TestUpdateCoinHandler(t *testing.T) {

	database.ConnectToDatabase()
	
	CreateCoin()
	defer DeleteCoin()
	
	r := SetUpRoutesTest()
	r.PATCH("/coin/:id", controllers.UpdateCoin)

	coin := models.Coin{Name: "CoinY", ShortName: "CNY", Votes: 2, Price: 2.50}
	coinJson, _ := json.Marshal(coin)
	routeTest := `/coin/` + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", routeTest, bytes.NewBuffer(coinJson))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var coinUpdated models.Coin
	json.Unmarshal(res.Body.Bytes(), &coinUpdated)


	assert.Equal(t, "CoinY", coinUpdated.Name, "The values must be the same")
	assert.Equal(t, "CNY", coinUpdated.ShortName, "The values must be the same")
	assert.Equal(t, 2, coinUpdated.Votes, "The values must be the same")
	assert.Equal(t, 2.50, coinUpdated.Price, "The values must be the same")
	assert.Equal(t, http.StatusOK, res.Code, "The status must be the same")
}