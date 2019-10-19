package presenters

import (
	"context"
	"net/http"

	"github.com/yuuis/RecommendSystem/api/utilities"
	"github.com/yuuis/RecommendSystem/models/recommend"
)

// RecommendView : レコメンドの返答
func RecommendView(ctx context.Context, r recommend.Recommend) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, r)
}
