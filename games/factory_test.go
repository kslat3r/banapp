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

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	data := gameData{
		Name: "test",
	}

	created, err := create(&data)

	assert.NoError(err)
	assert.IsType(created.ID, bson.NewObjectId())
	assert.Equal(created.Name, "test")
	assert.Equal(created.RemainingLetters, letters)
	assert.Equal(created.Status, statusWaiting)
	assert.Equal(created.Players, []player{})
}

func TestList(t *testing.T) {
	assert := assert.New(t)

	data := gameData{
		Name: "test",
	}

	create(&data)

	found, err := list()
	last := found[len(found)-1]

	assert.NoError(err)
	assert.IsType(last.ID, bson.NewObjectId())
	assert.Equal(last.Name, "test")
	assert.Equal(last.RemainingLetters, letters)
	assert.Equal(last.Status, statusWaiting)
	assert.Equal(last.Players, []player{})
}

func TestGetByID(t *testing.T) {
	assert := assert.New(t)

	data := gameData{
		Name: "test",
	}

	created, _ := create(&data)

	found, err := getByID(created.ID.Hex())

	assert.NoError(err)
	assert.IsType(found.ID, bson.NewObjectId())
	assert.Equal(found.Name, "test")
	assert.Equal(found.RemainingLetters, letters)
	assert.Equal(found.Status, statusWaiting)
	assert.Equal(found.Players, []player{})
}
