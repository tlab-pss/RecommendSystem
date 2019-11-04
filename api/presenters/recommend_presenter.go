package presenters

import (
	"context"
	"net/http"

	"github.com/yuuis/RecommendSystem/api/utilities"
	"github.com/yuuis/RecommendSystem/models/response"
)

// RecommendView : レコメンドの返答
func RecommendView(ctx context.Context, r response.ResponseType) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, r)
}
