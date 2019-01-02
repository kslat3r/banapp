package test

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// ListUsers List all users
func ListUsers(router *gin.RouterGroup) {
	router.GET("/users", func(c *gin.Context) {
		usersCh := make(chan []user, 1)

		go getUsers(usersCh)

		users := <-usersCh

		var wg sync.WaitGroup
		wg.Add(len(users))

		for i := range users {
			go users[i].GetTodos(&wg)
		}

		wg.Wait()

		c.JSON(http.StatusOK, users)
	})
}
