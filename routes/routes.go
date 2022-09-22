package routes

import (
	"github.com/CaioLuColaco/api-upVote-golang/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/coins", controllers.ShowCoins)
	r.Run(":2005")
}