package games

import (
	"banapp/common"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
)

func init() {
	common.InitDb()
}

func TestCreateGameRoute(t *testing.T) {
	assert := assert.New(t)

	engine := gin.New()
	api := engine.Group("/api")

	CreateGame(api.Group("/games"))

	g, _ := json.Marshal(game{
		Name: "test",
	})

	req, err := http.NewRequest("POST", "/api/games/", bytes.NewBuffer(g))
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)

	created := game{}
	json.Unmarshal([]byte(res.Body.String()), &created)

	assert.NoError(err)
	assert.IsType(created.ID, bson.NewObjectId())
	assert.Equal(created.Name, "test")
	assert.Equal(created.RemainingLetters, letters)
	assert.Equal(created.Status, statusWaiting)
	assert.Equal(created.Players, []player{})
}

func TestGetGamesRoute(t *testing.T) {
	assert := assert.New(t)

	data := gameData{
		Name: "test",
	}

	create(&data)

	engine := gin.New()
	api := engine.Group("/api")

	GetGames(api.Group("/games"))

	req, err := http.NewRequest("GET", "/api/games/", nil)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)

	list := []game{}
	json.Unmarshal([]byte(res.Body.String()), &list)

	last := list[len(list)-1]

	assert.NoError(err)
	assert.IsType(last.ID, bson.NewObjectId())
	assert.Equal(last.Name, "test")
	assert.Equal(last.RemainingLetters, letters)
	assert.Equal(last.Status, statusWaiting)
	assert.Equal(last.Players, []player{})
}

func TestGetGameRoute(t *testing.T) {
	assert := assert.New(t)

	data := gameData{
		Name: "test",
	}

	created, _ := create(&data)

	engine := gin.New()
	api := engine.Group("/api")

	GetGame(api.Group("/games"))

	req, err := http.NewRequest("GET", fmt.Sprintf("/api/games/%s", created.ID.Hex()), nil)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)

	found := game{}
	json.Unmarshal([]byte(res.Body.String()), &found)

	assert.NoError(err)
	assert.IsType(found.ID, bson.NewObjectId())
	assert.Equal(found.Name, "test")
	assert.Equal(found.RemainingLetters, letters)
	assert.Equal(found.Status, statusWaiting)
	assert.Equal(found.Players, []player{})
}

func TestUpdateRoute(t *testing.T) {

}

func TestAddPlayerRoute(t *testing.T) {

}
