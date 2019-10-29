package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/yuuis/RecommendSystem/api/presenters"
	"github.com/yuuis/RecommendSystem/api/utilities"
	"github.com/yuuis/RecommendSystem/models/recommend"
	"github.com/yuuis/RecommendSystem/models/request"
	"github.com/yuuis/RecommendSystem/models/service"
	"github.com/yuuis/RecommendSystem/modules/hotpepper"

	"github.com/gin-gonic/gin"
)

// Recommend : PAからレコメンドのリクエスト
func Recommend(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	var rrt request.ReceiveRequestType
	if err := json.Unmarshal(bodyBytes, &rrt); err != nil {
		presenters.ViewBadRequest(ctx, err)
	}

	// Note : プラグインサービスの選定
	plugin, err := service.SelectServicePlugin(rrt.TopicCategory)
	if err != nil {
		fmt.Printf("Plugin not found: %+v", rrt.TopicCategory)
		presenters.RecommendView(ctx, recommend.Recommend{Success: false})
	}

	// Todo : プラグイン情報から、外部サービスにリクエストをする
	fmt.Printf("Execute plugin: %+v \n", plugin)

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// todo: 現在は確定でhotpepperなので、とりあえず。
	if plugin.Name == "Hotpepper" {
		var hotpepperRrt hotpepper.ReceiveRequestType

		if err := c.BindJSON(&hotpepperRrt); err != nil {
			presenters.RecommendView(ctx, recommendFailure("error"))
		}

		rp := hotpepperRrt.RequestParameter

		shops, err := hotpepper.Request(&rp)

		if err != nil {
			presenters.RecommendView(ctx, recommendFailure(err.Error()))
			return
		}

		presenters.RecommendView(ctx, recommend.Recommend{
			Success:    true,
			Text:       shops[0].Name, // todo: とりあえず1個目
			ImagePaths: nil,
		})
	} else {
		presenters.RecommendView(ctx, recommend.Recommend{
			Success:    true,
			Text:       "いまのところhotpepper以外はないよん",
			ImagePaths: nil,
		})
	}
}

func recommendFailure(text string) recommend.Recommend {
	return recommend.Recommend{
		Success:    false,
		Text:       text,
		ImagePaths: nil,
	}
}
