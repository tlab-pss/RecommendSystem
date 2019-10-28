package api

import (
	"github.com/yuuis/RecommendSystem/api/controllers"

	"github.com/gin-gonic/gin"
)

func Router(e *gin.Engine) {
	api := e.Group("/api")

	// PAからサービスのリクエストを受ける
	api.POST("/recommend", controllers.Recommend)

}
