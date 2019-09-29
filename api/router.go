package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yuuis/RecommendSystem/api/controllers"
)

func Router(e *gin.Engine) {
	api := e.Group("/api")

	// todo: サンプル実装。最初のprで下の1行を削除
	api.GET("/hello", controllers.Hello)
}
