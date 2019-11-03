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
	"github.com/yuuis/RecommendSystem/models/response"
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
		return
	}

	var rrt request.ReceiveRequestType
	if err := json.Unmarshal(bodyBytes, &rrt); err != nil {
		presenters.ViewBadRequest(ctx, err)
		return
	}

	fmt.Printf("Params: %+v", rrt)

	// Note : プラグインサービスの選定
	plugin, err := service.SelectServicePlugin(rrt.TopicCategory)
	if err != nil {
		fmt.Printf("Plugin not found: %+v", rrt.TopicCategory)

		presenters.RecommendView(ctx, response.ResponseType{
			ResponseData: recommend.Recommend{
				Success: false,
			},
		})
		return
	}

	// Todo : プラグイン情報から、外部サービスにリクエストをする
	fmt.Printf("Execute plugin: %+v \n", plugin)

	responseValue := &response.ResponseType{
		UUID:        plugin.ID, // TODO : リクエストの履歴をDBに保存したい
		ServiceName: plugin.Name,
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// todo: 現在は確定でhotpepperなので、とりあえず。
	if plugin.Name == "hotpepper" {
		var hotpepperRrt hotpepper.ReceiveRequestType

		if err := c.BindJSON(&hotpepperRrt); err != nil {
			responseValue.ResponseData = recommendFailure("error")
			goto Return
		}

		rp := hotpepperRrt.RequestParameter

		shops, err := hotpepper.Request(&rp)
		if err != nil {
			responseValue.ResponseData = recommendFailure(err.Error())
			goto Return
		}

		responseValue.ResponseData = recommend.Recommend{
			Success:    true,
			Text:       shops[0].Name, // todo: とりあえず1個目
			ImagePaths: nil,
		}
	} else {
		responseValue.ResponseData = recommend.Recommend{
			Success:    true,
			Text:       "いまのところhotpepper以外はないよん",
			ImagePaths: nil,
		}
	}

Return:
	presenters.RecommendView(ctx, *responseValue)
}

func recommendFailure(text string) recommend.Recommend {
	return recommend.Recommend{
		Success:    false,
		Text:       text,
		ImagePaths: nil,
	}
}
