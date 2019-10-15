package controllers

import (
	"github.com/gin-gonic/gin"
	"main/api/presenters"
	"main/api/utilities"
	"main/models/hello"
)

// todo: サンプル実装。最初のprでこのファイルを削除
func Hello(c *gin.Context) {
	ctx := utilities.AddGinContext(c.Request.Context(), c)

	presenters.HelloView(ctx, hello.Hello{Message: "hello"})
}
