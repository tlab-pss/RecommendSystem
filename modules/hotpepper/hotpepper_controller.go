package hotpepper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Request : APIにリクエストを投げる
func Request() (*Payload, error) {
	replyData := new(Payload)

	req, err := http.NewRequest("GET", "http://pd:8080/api/plugin-services", nil)
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
