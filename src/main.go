package main

import (
	"github.com/gin-gonic/gin"

	"bananagrams-api/games"
)

func main() {
	engine := gin.Default()
	api := engine.Group("/api")

	games.List(api.Group("/games"))
	games.Create(api.Group("/games"))

	engine.Run()
}
