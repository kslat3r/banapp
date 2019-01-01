package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// List all games
func List(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, list())
	})

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "hello")
	})
}

// Create game
func Create(router *gin.RouterGroup) {
	router.POST("/", func(c *gin.Context) {
		var g game
		c.BindJSON(&g)

		g = create(g)

		c.JSON(http.StatusOK, g)
	})
}
