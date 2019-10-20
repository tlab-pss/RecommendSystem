package hotpepper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Request : APIにリクエストを投げる
func Request(payload *Payload) (*ResponseType, error) {
	replyData := new(ResponseType)

	req, err := http.NewRequest("GET", "http://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key="+os.Getenv("RECRUIT_APIKEY")+"&keyword="+payload.Keywords, nil)
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

	if err := json.NewDecoder(r).Decode(replyData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}

	return replyData, nil
}
