package serviceselector

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// PluginService : サービスプラグインの型
type PluginService struct {
	ID            string
	Name          string
	BigCategoryID string
	CreatedAt     time.Time
}

// GetPluginService : プラグインされたサービスを取得
func GetPluginService() (*[]PluginService, error) {
	replyData := new([]PluginService)

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

// SelectService : プラグインさせれたサービスを
func SelectService() {

}
