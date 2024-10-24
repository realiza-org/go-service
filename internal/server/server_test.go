package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	// Initialize the server
	e := NewServer()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assert that the server initializes correctly
	e.ServeHTTP(c.Response(), c.Request())
	if assert.NoError(t, nil) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
