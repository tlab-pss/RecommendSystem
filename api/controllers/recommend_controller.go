package controllers

import (
	"fmt"

	"github.com/yuuis/RecommendSystem/api/presenters"
	"github.com/yuuis/RecommendSystem/api/utilities"
	"github.com/yuuis/RecommendSystem/models/recommend"
	"github.com/yuuis/RecommendSystem/models/service"

	"github.com/gin-gonic/gin"
)

// ReceiveRequestType : 送られてきたリクエストに付随する値を格納する型
type ReceiveRequestType struct {
	TopicCategory    service.ServiceCategory `json:"topic_category"`
	RequireService   bool                    `json:"require_service"`
	ServiceDataValue interface{}             `json:"service_data_value"`
}

// Recommend : PAからレコメンドのリクエスト
func Recommend(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)

	var rrt ReceiveRequestType
	if err := c.BindJSON(&rrt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	// Note : プラグインサービスの選定
	plugin, err := service.PluginServiceSelector(rrt.TopicCategory)
	if err != nil {
		fmt.Printf("Plugin not found: %+v", rrt.TopicCategory)
		presenters.RecommendView(ctx, recommend.Recommend{Success: false})
	}

	// Todo : プラグイン情報から、外部サービスにリクエストをする
	fmt.Printf("Execute plugin: %+v", plugin)

	presenters.RecommendView(ctx, recommend.Recommend{Success: true, Text: plugin.Name})
}
