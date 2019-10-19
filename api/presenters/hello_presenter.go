package presenters

import (
	"context"
	"net/http"

	"github.com/yuuis/RecommendSystem/api/utilities"
	"github.com/yuuis/RecommendSystem/models/hello"
)

// todo: サンプル実装。最初のprでこのを削除
func HelloView(ctx context.Context, h hello.Hello) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, h)
}
