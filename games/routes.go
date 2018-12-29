package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, getAll())
	})
}

func Create(router *gin.RouterGroup) {
	router.POST("/", func(c *gin.Context) {
		var newGame game
		c.BindJSON(&newGame)

		create(newGame)

		c.JSON(http.StatusOK, newGame)
	})
}
