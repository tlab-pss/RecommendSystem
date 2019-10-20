package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuis/RecommendSystem/api"
	"github.com/yuuis/RecommendSystem/infrastructures"
)

func TestHelloSuccess(t *testing.T) {

	s := infrastructures.NewServer()
	api.Router(s)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/hello", nil)
	s.ServeHTTP(w, req)

	json := `{"Message":"hello"}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, json, w.Body.String())
}
