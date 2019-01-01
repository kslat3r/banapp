package games

import (
	"ban-app/common"
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
)

func init() {
	common.InitDb()
}

func TestList(t *testing.T) {
	assert := assert.New(t)
	games := list()

	for _, game := range games {
		assert.IsType(game.ID, bson.NewObjectId())
		assert.NotEmpty(game.Name)
	}
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	newGame := game{
		Name: "foo",
	}

	newGame = create(newGame)

	assert.IsType(newGame.ID, bson.NewObjectId())
	assert.Equal(newGame.Name, "foo")
}
