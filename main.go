package main

import (
	"github.com/gin-gonic/gin"

	"ban-app/common"
	"ban-app/games"
	"ban-app/test"
)

func init() {
	common.InitDb()
}

func main() {
	router := gin.Default()
	api := router.Group("/api")

	gamesGroup := api.Group("/games")

	games.List(gamesGroup)
	games.Create(gamesGroup)

	testGroup := api.Group("/test")

	test.ListUsers(testGroup)

	router.Run()
}
