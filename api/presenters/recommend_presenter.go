package presenters

import (
	"context"
	"main/api/utilities"
	"main/models/recommend"
	"net/http"
)

// RecommendView : レコメンドの返答
func RecommendView(ctx context.Context, r recommend.Recommend) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusOK, r)
}
