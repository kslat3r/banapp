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
