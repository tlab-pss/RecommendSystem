package hotpepper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Request : hotpepper APIにリクエストを投げる
func Request(parameter *RequestParameter) ([]Shop, error) {

	req, err := http.NewRequest("GET", "http://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key="+os.Getenv("RECRUIT_APIKEY")+"&keyword="+parameter.Keywords, nil)
	if err != nil {
		fmt.Printf("pd error, cannot create http request")
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("pd error! cannot exec http request")
		return nil, err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	// r = io.TeeReader(r, os.Stderr)

	hr := hotpepperResponse{}
	if err := json.NewDecoder(r).Decode(hr); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}

	return hr.Results.Shop, nil
}

// RequestParameter : リクエストパラメタ
type RequestParameter struct {
	Keywords string `json:"keywords"`
}

// HotpepperResponse: APIResponseの型
type hotpepperResponse struct {
	Results struct {
		ResultsStart     int    `json:"results_start"`
		ResultsReturned  string `json:"results_returned"`
		APIVersion       string `json:"api_version"`
		Shop             []Shop `json:"shop"`
		ResultsAvailable int    `json:"results_available"`
		Error            []struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	} `json:"results"`
}
