package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/RecommendSystem/api/presenters"
	"github.com/yuuis/RecommendSystem/api/utilities"
	"github.com/yuuis/RecommendSystem/models/hello"
)

// todo: サンプル実装。最初のprでこのファイルを削除
func Hello(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)

	presenters.HelloView(ctx, hello.Hello{Message: "hello"})
}
