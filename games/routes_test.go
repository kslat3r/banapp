package games

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	assert := assert.New(t)

	engine := gin.New()
	api := engine.Group("/api")

	List(api.Group("/games"))

	req, err := http.NewRequest("GET", "/api/games", nil)
	assert.NoError(err)

	recorder := httptest.NewRecorder()
	engine.ServeHTTP(recorder, req)

	result := []game{}
	json.NewDecoder(recorder.Body).Decode(&result)

	for _, game := range result {
		assert.NotNil(game.ID)
	}
}
