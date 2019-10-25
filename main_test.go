package main

import (
	"bytes"
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

func TestServiceRequestToHotpepper(t *testing.T) {
	infrastructures.InitEnvironment()

	s := infrastructures.NewServer()
	api.Router(s)

	requestData := `
	{
		"topic_category": 2,
		"require_service": true,
		"service_data_value": {
			"keywords": "寿司"
		}
	}`

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/recommend", bytes.NewBuffer([]byte(requestData)))
	s.ServeHTTP(w, req)

	json := `{"success":true,"text":"いろり屋 iroriya 新橋駅前店","image_paths":null}`
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, json, w.Body.String())
}
