package controllers

import (
	"fmt"
	"main/api/presenters"
	"main/api/utilities"
	"main/models/recommend"

	"github.com/gin-gonic/gin"

	serviceselector "main/modules/selector"
)

// Recommend : PAからレコメンドのリクエスト
func Recommend(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)

	var rrt recommend.ReceiveRequestType
	if err := c.BindJSON(&rrt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	// Note : プラグインサービスの選定
	plugin, err := serviceselector.PluginServiceSelector(rrt.TopicCategory)
	if err != nil {
		fmt.Printf("Plugin not found: %+v", rrt.TopicCategory)
		presenters.RecommendView(ctx, recommend.Recommend{Success: false})
	}

	// Todo : プラグイン情報から、外部サービスにリクエストをする
	fmt.Printf("Execute plugin: %+v", plugin)

	presenters.RecommendView(ctx, recommend.Recommend{Success: true, Text: plugin.Name})
}
