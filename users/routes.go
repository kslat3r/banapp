package users

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Get List all users
func Get(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		users := get()

		wg := sync.WaitGroup{}
		wg.Add(len(users))

		for i := range users {
			go users[i].getTodos(&wg)
		}

		wg.Wait()
		c.JSON(http.StatusOK, users)
	})
}
