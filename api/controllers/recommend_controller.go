package controllers

import (
	"main/api/presenters"
	"main/api/utilities"
	"main/models/recommend"

	"github.com/gin-gonic/gin"
)

// Recommend : PDからレコメンドのリクエスト
func Recommend(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)

	// Todo : プラグインサービスの選定

	presenters.RecommendView(ctx, recommend.Recommend{Text: "hello"})
}
