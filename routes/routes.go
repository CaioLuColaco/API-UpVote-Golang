package routes

import (
	"github.com/CaioLuColaco/api-upVote-golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/coins", controllers.ShowAllCoins)
	r.GET("/coins/:id", controllers.ShowOneCoin)
	r.POST("/coin", controllers.CreateCoin)
	r.DELETE("/coin/:id", controllers.DeleteCoin)
	r.PATCH("/coin/:id", controllers.UpdateCoin)
	r.Run(":2005")
}