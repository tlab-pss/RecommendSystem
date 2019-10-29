package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuuis/RecommendSystem/api"
	"github.com/yuuis/RecommendSystem/infrastructures"
	"github.com/yuuis/RecommendSystem/models/recommend"
)

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

	assert.Equal(t, 200, w.Code)

	var res recommend.Recommend
	if err := json.Unmarshal(w.Body.Bytes(), &res); err != nil {
		fmt.Println(fmt.Errorf("Response json parse error: %+v", err))
	}
	assert.Equal(t, true, res.Success)
}
