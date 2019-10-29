package hotpepper

import (
	"github.com/yuuis/RecommendSystem/models/basic_location"
	"github.com/yuuis/RecommendSystem/models/location"
	"math"
)

func Optimize() (*Condition, error) {
	// todo: 位置情報とか食べたものとかから、hotpepper用のconditionを作って返すぞ！

	// 位置
	bl, err := basic_location.Get()
	if err != nil {
		return &Condition{}, err
	}

	l, err := location.GetLatest()
	if err != nil {
		return &Condition{}, err
	}

	// 現在地と家との差
	hlatDiff :=math.Abs(bl.House.Latitude - l.Latitude)
	hlngDiff := math.Abs(bl.House.Longitude - l.Longitude)

	// 現在地とオフィスとの差
	olatDiff :=math.Abs(bl.Office.Latitude - l.Latitude)
	olngDiff := math.Abs(bl.Office.Longitude - l.Longitude)

	// todo: 位置の判定が適当すぎる
	if hlatDiff < 0.001 && hlngDiff < 0.001 {
		// todo: 家にいる人へのconditionつくる
	} else if olatDiff < 0.001 && olngDiff < 0.001 {
		// todo: オフィスにいる人へのconditionつくる
	} else {
		// todo: 外にいる人へのconditionつくる
	}


	// 食べたもの
  // todo: 昨日食べたものとかからよしなに

	return &Condition{}, nil
}

type Condition struct {
	Condition []map[string]interface{}
}
