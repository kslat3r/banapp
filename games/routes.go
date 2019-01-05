package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateGame create a game
func CreateGame(router *gin.RouterGroup) {
	router.POST("/", func(c *gin.Context) {
		g := game{}
		err := c.BindJSON(&g)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, createGame(&g))
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
		err := c.BindJSON(&ug)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, updateGame(&g, &ug))
	})
}

// CreatePlayer add player to game
func CreatePlayer(router *gin.RouterGroup) {
	router.POST("/:id/players", func(c *gin.Context) {
		g := getGame(c.Param("id"))
		p := player{}
		err := c.BindJSON(&p)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, addPlayer(&g, &p))
	})
}
