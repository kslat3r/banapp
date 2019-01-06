package games

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type gameData struct {
	Name   string `json:"name"`
	Status status `json:"status"`
}

type playerData struct {
	Name string `json:"name"`
}

// CreateGame create a game
func CreateGame(router *gin.RouterGroup) {
	router.POST("/", func(c *gin.Context) {
		data := gameData{}
		err := c.BindJSON(&data)

		if err != nil {
			panic(err)
		}

		game, err := create(&data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, game)
	})
}

// GetGames get all games
func GetGames(router *gin.RouterGroup) {
	router.GET("/", func(c *gin.Context) {
		found, err := list()

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, found)
	})
}

// GetGame get a game
func GetGame(router *gin.RouterGroup) {
	router.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		found, err := getByID(id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, found)
	})
}

// UpdateGame update a game
func UpdateGame(router *gin.RouterGroup) {
	router.PATCH("/:id", func(c *gin.Context) {
		id := c.Param("id")
		game, err := getByID(id)

		if err != nil {
			panic(err)
		}

		data := gameData{}
		err = c.BindJSON(&data)

		if err != nil {
			panic(err)
		}

		_, err = game.update(&data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, game)
	})
}

// AddPlayer add player to game
func AddPlayer(router *gin.RouterGroup) {
	router.POST("/:id/players", func(c *gin.Context) {
		id := c.Param("id")
		game, err := getByID(id)

		if err != nil {
			panic(err)
		}

		data := playerData{}
		err = c.BindJSON(&data)

		if err != nil {
			panic(err)
		}

		_, err = game.addPlayer(&data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, game)
	})
}
