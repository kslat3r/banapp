package games

import (
	"banapp/common"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
)

func init() {
	common.InitDb()
}

func TestCreateGame(t *testing.T) {
	assert := assert.New(t)
	newGame := game{
		Name: "foo",
	}

	newGame = createGame(&newGame)

	assert.IsType(newGame.ID, bson.NewObjectId())
	assert.Equal(newGame.Name, "foo")
}

func TestGetGames(t *testing.T) {
	assert := assert.New(t)
	games := getGames()

	for _, game := range games {
		assert.IsType(game.ID, bson.NewObjectId())
		assert.NotEmpty(game.Name)
	}
}
