package main

import (
	"github.com/gin-gonic/gin"

	"ban-api/src/common"
	"ban-api/src/games"
)

func init() {
	common.InitDb()
}

func main() {
	router := gin.Default()
	api := router.Group("/api")

	games.List(api.Group("/games"))
	games.Create(api.Group("/games"))

	router.Run()
}
