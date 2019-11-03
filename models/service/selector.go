package service

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// GetAllPluginService : プラグインされたサービスを取得
func GetAllPluginService() ([]PluginService, error) {
	var replyData []PluginService

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

	if err := json.NewDecoder(r).Decode(&replyData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}

	return replyData, nil
}

// SelectServicePlugin : プラグインサービスを選別する関数
func SelectServicePlugin(sc ServiceCategory) (*PluginService, error) {

	var pluginService PluginService

	allServices, err := GetAllPluginService()
	if err != nil {
		return &pluginService, err
	}

	var pluginList []PluginService
	for _, m := range allServices {
		pluginList = append(pluginList, m)
	}

	// TODO : 今はランダム。利用傾向や満足度などからよしなにしたい
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(pluginList))
	pluginService = pluginList[i]

	return &pluginService, nil
}
