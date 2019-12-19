package server

import (
	"quote-me/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter function initialises the backend
func InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("api")
	{
		quoteController := new(controllers.QuoteController)
		api.GET("/quote", quoteController.Quote)
	}

	return router
}
