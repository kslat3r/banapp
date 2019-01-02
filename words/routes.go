package words

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get List all users
func Get(router *gin.RouterGroup) {
	router.GET("/:word", func(c *gin.Context) {
		c.JSON(http.StatusOK, get(c.Param("word")))
	})
}
