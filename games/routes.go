package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateGame create a game
func CreateGame(router *gin.RouterGroup) {
	router.POST("/", func(c *gin.Context) {
		g := game{}
		c.BindJSON(&g)

		c.JSON(http.StatusOK, createGame(g))
	})
}

// GetGames get all games
func GetGames(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, getGames())
	})
}

// GetGame get a game
func GetGame(router *gin.RouterGroup) {
	router.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, getGame(c.Param("id")))
	})
}

// UpdateGame update a game
func UpdateGame(router *gin.RouterGroup) {
	router.PATCH("/:id", func(c *gin.Context) {
		g := getGame(c.Param("id"))

		ug := game{}
		c.BindJSON(&ug)

		c.JSON(http.StatusOK, updateGame(g.ID, ug))
	})
}

// CreatePlayer add player to game
func CreatePlayer(router *gin.RouterGroup) {
	router.POST("/:id/players", func(c *gin.Context) {
		g := getGame(c.Param("id"))

		p := player{}
		c.BindJSON(&p)

		c.JSON(http.StatusOK, createPlayer(g.ID, p))
	})
}
