package games

import (
	"ban-app/common"
	"bytes"
	"encoding/json"
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

func TestGetRoute(t *testing.T) {
	assert := assert.New(t)

	engine := gin.New()
	api := engine.Group("/api")

	Get(api.Group("/games"))

	req, err := http.NewRequest("GET", "/api/games/", nil)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)

	list := []game{}
	json.Unmarshal([]byte(res.Body.String()), &list)

	assert.NoError(err)
	assert.IsType(list[len(list)-1].ID, bson.NewObjectId())
	assert.NotEmpty(list[len(list)-1].Name)
}

func TestCreateRoute(t *testing.T) {
	assert := assert.New(t)

	engine := gin.New()
	api := engine.Group("/api")

	Create(api.Group("/games"))

	g, _ := json.Marshal(game{
		Name: "bar",
	})

	req, err := http.NewRequest("POST", "/api/games/", bytes.NewBuffer(g))
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)

	created := game{}
	json.Unmarshal([]byte(res.Body.String()), &created)

	assert.NoError(err)
	assert.IsType(created.ID, bson.NewObjectId())
	assert.Equal(created.Name, "bar")
}
