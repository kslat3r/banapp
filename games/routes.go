package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all games
func Get(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, get())
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
