package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func filterTodos(todos []todo, f func(todo) bool) []todo {
	filtered := make([]todo, 0)

	for _, todo := range todos {
		if f(todo) {
			filtered = append(filtered, todo)
		}
	}

	return filtered
}

// ListUsers List all users
func ListUsers(router *gin.RouterGroup) {
	router.GET("/users", func(c *gin.Context) {
		usersCh := make(chan []user)
		todosCh := make(chan []todo)

		go listUsers(usersCh)
		go listTodos(todosCh)

		users := <-usersCh
		todos := <-todosCh

		for i := range users {
			users[i].Todos = filterTodos(todos, func(todo todo) bool {
				return todo.UserID == users[i].ID
			})
		}

		c.JSON(http.StatusOK, users)
	})
}
