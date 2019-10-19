package service

import (
	"time"
)

// PluginService : サービスプラグインの型
type PluginService struct {
	ID            string
	Name          string
	BigCategoryID string
	CreatedAt     time.Time
}

func (p PluginService) toServiceCategory() ServiceCategory {

	bigCategoryID := p.BigCategoryID

	switch bigCategoryID {
	case "0":
		return Commerce
	case "1":
		return Gourmet
	case "2":
		return Weather
	case "3":
		return Map
	case "4":
		return Mail
	case "5":
		return Music
	case "6":
		return Message
	case "7":
		return Search
	case "8":
		return Translation
	case "9":
		return News
	default:
		return Uncategorized
	}
}
