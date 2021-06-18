package router

import (
	"htmlparser/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, "Page not found")
	})

	group := router.Group("/parse")
	{
		group.POST("/url", controllers.ParseHandler)
	}

	return router
}
