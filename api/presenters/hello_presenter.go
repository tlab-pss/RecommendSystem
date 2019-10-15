package presenters

import (
	"context"
	"main/api/utilities"
	"main/models/hello"
	"net/http"
)

// todo: サンプル実装。最初のprでこのを削除
func HelloView(ctx context.Context, h hello.Hello) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, h)
}
