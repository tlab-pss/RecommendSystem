package presenters

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yuuis/RecommendSystem/api/utilities"
	"net/http"
)

func ViewInternalServerError(ctx context.Context, err error) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusInternalServerError, map[string]interface{}{"errors": err.Error()})
}

func ViewBadRequest(ctx context.Context, err error) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	JSON(c, http.StatusBadRequest, map[string]interface{}{"errors": err.Error()})
}

func ViewNoContent(ctx context.Context) {
	defer utilities.DeleteGinContext(ctx)
	c := utilities.GetGinContext(ctx)
	c.Writer.WriteHeader(http.StatusNoContent)
}

func JSON(c *gin.Context, code int, v interface{}) {
	c.JSON(code, v)
}
