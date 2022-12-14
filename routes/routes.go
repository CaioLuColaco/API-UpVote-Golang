package routes

import (
	"github.com/CaioLuColaco/api-upVote-golang/controllers"
	"github.com/gin-contrib/cors"
	docs "github.com/CaioLuColaco/api-upVote-golang/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/gin-gonic/gin"
	"os"
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

	// r.Run("0.0.0.0:2005")
	if os.Getenv("PORT") != "" {
		// Heroku add a env variable called PORT, if exist we will use it
		r.Run("0.0.0.0:" + os.Getenv("PORT"))
	} else {
		// If is running on localhost (our computer), no PORT env variable
		r.Run("0.0.0.0:8080")
	}
}