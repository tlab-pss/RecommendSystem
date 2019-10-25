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
	plugin, err := service.PluginServiceSelector(rrt.TopicCategory)
	if err != nil {
		fmt.Printf("Plugin not found: %+v", rrt.TopicCategory)
		presenters.RecommendView(ctx, recommend.Recommend{Success: false})
	}

	// Todo : プラグイン情報から、外部サービスにリクエストをする
	fmt.Printf("Execute plugin: %+v \n", plugin)

	result := new(recommend.Recommend)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if plugin.Name == "Hotpepper" {
		var hotpepperRrt hotpepper.ReceiveRequestType
		var payload hotpepper.Payload
		if err := c.BindJSON(&hotpepperRrt); err != nil {
			payload = hotpepper.Payload{
				Keywords: "焼肉",
			}
			result.Success = false
			result.Text = "error"
			presenters.RecommendView(ctx, *result)
		} else {
			payload = hotpepperRrt.ServiceDataValue
		}

		serviceResponse, err := hotpepper.Request(&payload)

		if err != nil {
			result.Success = false
			result.Text = err.Error()
			presenters.RecommendView(ctx, *result)
		}

		shops, err := serviceResponse.GetShopNames()
		if err != nil {
			result.Success = false
			result.Text = err.Error()
			presenters.RecommendView(ctx, *result)
		}
		result.Text = shops[0]
	}

	result.Success = true
	presenters.RecommendView(ctx, *result)
}
