package games

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestListRoute(t *testing.T) {
	assert := assert.New(t)

	engine := gin.New()
	api := engine.Group("/api")

	List(api.Group("/games"))

	req, err := http.NewRequest("GET", "/api/games/", nil)
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)

	list := []game{}
	json.Unmarshal([]byte(res.Body.String()), &list)

	assert.NoError(err)
	assert.NotEmpty(list[len(list)-1].ID)
	assert.NotEmpty(list[len(list)-1].Name)
}

func TestCreateRoute(t *testing.T) {
	assert := assert.New(t)

	engine := gin.New()
	api := engine.Group("/api")

	Create(api.Group("/games"))

	g, _ := json.Marshal(game{
		ID:   "foo",
		Name: "bar",
	})

	req, err := http.NewRequest("POST", "/api/games/", bytes.NewBuffer(g))
	res := httptest.NewRecorder()
	engine.ServeHTTP(res, req)

	created := game{}
	json.Unmarshal([]byte(res.Body.String()), &created)

	assert.NoError(err)
	assert.NotEmpty(created.ID)
	assert.NotEqual(created.ID, "foo")
	assert.Equal(created.Name, "bar")
}
