package routes

import (
	"github.com/CaioLuColaco/api-upVote-golang/controllers"
	"github.com/gin-contrib/cors"
	docs "github.com/CaioLuColaco/api-upVote-golang/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.Use(cors.Default())
	r.GET("/coins", controllers.ShowAllCoins)
	r.GET("/coins/:id", controllers.ShowOneCoin)
	r.GET("/coins/name/:name", controllers.ShowCoinByName)
	r.GET("/coins/shortName/:shortname", controllers.ShowCoinByShortName)

	r.POST("/coin", controllers.CreateCoin)
	r.PATCH("/coin/:id", controllers.UpdateCoin)
	
	r.DELETE("/coin/:id", controllers.DeleteCoin)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run("0.0.0.0:2005")
}