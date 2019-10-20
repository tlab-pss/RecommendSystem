package api

import (
	"github.com/yuuis/RecommendSystem/api/controllers"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	api := e.Group("/api")

	// todo: サンプル実装。最初のprで下の1行を削除
	api.GET("/hello", controllers.Hello)

	// PAからサービスのリクエストを受ける
	api.GET("/request", controllers.Recommend)

}
