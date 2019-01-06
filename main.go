package main

import (
	"github.com/gin-gonic/gin"

	"banapp/common"
	"banapp/games"
	"banapp/words"
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
	games.AddPlayer(gamesGroup)

	wordsGroup := api.Group("/words")
	words.Get(wordsGroup)

	router.Run()
}
