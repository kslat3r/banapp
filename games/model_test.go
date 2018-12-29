package games

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	assert := assert.New(t)
	games := list()

	assert.Equal(len(games), len(all))

	for key, game := range games {
		assert.Equal(game.ID, all[key].ID)
	}
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	newGame := game{
		Name: "name",
	}

	create(newGame)

	allGames := list()
	gameToTest := allGames[len(allGames)-1]

	assert.IsType(gameToTest.ID, "id")
	assert.Equal(gameToTest.Name, newGame.Name)
}
