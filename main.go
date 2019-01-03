package main

import (
	"github.com/gin-gonic/gin"

	"ban-app/common"
	"ban-app/games"
	"ban-app/users"
	"ban-app/words"
)

func init() {
	common.InitDb()
}

func main() {
	router := gin.Default()
	api := router.Group("/api")

	gamesGroup := api.Group("/games")
	games.CreateGame(gamesGroup)
	games.GetGames(gamesGroup)
	games.GetGame(gamesGroup)
	games.UpdateGame(gamesGroup)
	games.CreatePlayer(gamesGroup)

	usersGroup := api.Group("/users")
	users.Get(usersGroup)

	wordsGroup := api.Group("/words")
	words.Get(wordsGroup)

	router.Run()
}
