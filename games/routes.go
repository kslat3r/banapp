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
}

// Create game
func Create(router *gin.RouterGroup) {
	router.POST("/", func(c *gin.Context) {
		g := game{}
		c.BindJSON(&g)

		c.JSON(http.StatusOK, create(g))
	})
}
