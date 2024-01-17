package routes

import (
	"github.com/gabrielebonfim/cat-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/cats", handlers.GetCats)
	router.GET("/cats/:id", handlers.GetCatByID)
	router.POST("/cats", handlers.PostCat)

	return router
}
