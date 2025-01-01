package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetPosts(t *testing.T) {
	req, err := http.NewRequest("GET", "/posts", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetPosts(nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestCreatePost(t *testing.T) {
	body := `{"title": "Test Post", "content": "This is a test post."}`
	req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer([]byte(body)))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreatePost(nil))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}
