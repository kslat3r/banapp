package main

import (
	"github.com/gin-gonic/gin"

	"ban-app/common"
	"ban-app/games"
	"ban-app/users"
)

func init() {
	common.InitDb()
}

func main() {
	router := gin.Default()
	api := router.Group("/api")

	gamesGroup := api.Group("/games")

	games.Get(gamesGroup)
	games.Create(gamesGroup)

	usersGroup := api.Group("/users")

	users.Get(usersGroup)

	router.Run()
}
