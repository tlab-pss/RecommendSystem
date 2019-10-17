package serviceselector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	funk "github.com/thoas/go-funk"

	"main/models/service"
)

// PluginService : サービスプラグインの型
type PluginService struct {
	ID            string
	Name          string
	BigCategoryID string
	CreatedAt     time.Time
}

func (p PluginService) toServiceCategory() service.ServiceCategory {

	bigCategoryID := p.BigCategoryID

	switch bigCategoryID {
	case "0":
		return service.Commerce
	case "1":
		return service.Gourmet
	case "2":
		return service.Weather
	case "3":
		return service.Map
	case "4":
		return service.Mail
	case "5":
		return service.Music
	case "6":
		return service.Message
	case "7":
		return service.Search
	case "8":
		return service.Translation
	case "9":
		return service.News
	default:
		return service.Uncategorized
	}
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

// PluginServiceSelector : プラグインサービスを選別する関数
func PluginServiceSelector(sc service.ServiceCategory) (*PluginService, error) {

	var pluginService PluginService

	pluginServices, err := GetPluginService()
	if err != nil {
		return &pluginService, err
	}

	pluginList, ok := funk.Filter(pluginServices, func(plugin *PluginService) bool {
		return plugin.toServiceCategory() == sc
	}).([]PluginService)
	if ok != true {
		return &pluginService, errors.New("Could not cast PluginService")
	}

	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(pluginList))
	pluginService = pluginList[i]

	return &pluginService, nil
}
