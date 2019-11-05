package service

import (
	"encoding/json"
	"fmt"
	"github.com/yuuis/RecommendSystem/models/basic_location"
	"github.com/yuuis/RecommendSystem/models/location"
	"io"
	"math"
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
func SelectServicePlugin(sc ServiceCategory, l *location.Location) (*PluginService, error) {

	var pluginService PluginService

	allServices, err := GetAllPluginService()
	if err != nil {
		return &pluginService, err
	}

	var pluginList []PluginService
	for _, m := range allServices {
		pluginList = append(pluginList, m)
	}

	// todo: 位置情報で分ける
	// ユーザの位置
	bl, err := basic_location.Get()
	if err != nil {
		return &pluginService, err
	}

	// 現在地と家との差
	hlatDiff := math.Abs(bl.House.Latitude - l.Latitude)
	hlngDiff := math.Abs(bl.House.Longitude - l.Longitude)

	// 現在地とオフィスとの差
	olatDiff := math.Abs(bl.Office.Latitude - l.Latitude)
	olngDiff := math.Abs(bl.Office.Longitude - l.Longitude)

	// todo: 位置の判定が適当すぎる
	if hlatDiff < 0.001 && hlngDiff < 0.001 {
		//  家にいる人
		return &PluginService{
			ID:            "delivery",
			Name:          "delivery",
			BigCategoryID: string(sc),
			CreatedAt:     time.Time{},
		}, nil
	} else if olatDiff < 0.001 && olngDiff < 0.001 {
		// オフィスにいる人
		return &PluginService{
			ID:            "hotpepper",
			Name:          "hotpepper",
			BigCategoryID: string(sc),
			CreatedAt:     time.Time{},
		}, nil
	} else {
		// 外にいる人
		return &PluginService{
			ID:            "hotpepper",
			Name:          "hotpepper",
			BigCategoryID: string(sc),
			CreatedAt:     time.Time{},
		}, nil
	}
}
