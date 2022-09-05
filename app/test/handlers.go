package test

import (
	"encoding/json"
	"memoria/app/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/search/song", controllers.SearchSongs)

	return router
}

func TestSpotifyHandler(t *testing.T) {
	body := gin.H{
		"ID":         "3AJwUDP919kvQ9QcozQPxg",
		"name":       "Yellow",
		"artist":     "Coldplay",
		"albumCover": "https://i.scdn.co/image/ab67616d0000b2733d92b2ad5af9fbc8637425f0",
	}
	router := Setup()
	w := performRequest(router, "GET", "/search/song?name=yellow")
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, exists := response["ID"]
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["ID"], value)
}
