package routes

import (
	"github.com/CaioLuColaco/api-upVote-golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/coins", controllers.ShowAllCoins)
	r.GET("/coins/:id", controllers.ShowOneCoin)
	r.GET("/coins/name/:name", controllers.ShowCoinByName)
	r.GET("/coins/shortName/:shortname", controllers.ShowCoinByShortName)

	r.POST("/coin", controllers.CreateCoin)
	r.PATCH("/coin/:id", controllers.UpdateCoin)

	r.DELETE("/coin/:id", controllers.DeleteCoin)
	
	r.Run(":2005")
}