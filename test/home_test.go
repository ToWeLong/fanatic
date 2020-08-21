package test

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	app := setupApp()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	bytes := w.Body.Bytes()
	author := jsoniter.Get(bytes, "author")
	assert.Equal(t, "welong", author.ToString())
}
