package basic_location

import (
	"github.com/yuuis/RecommendSystem/models/location"
)

type BasicLocation struct {
	ID     string            `json:"ID"`
	House  location.Location `json:"House"`
	Office location.Location `json:"Office"`
	//Route     []location.Location `json:"Route"` todo: 未実装
}
