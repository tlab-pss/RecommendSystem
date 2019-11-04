package response

import (
	"github.com/yuuis/RecommendSystem/models/recommend"
)

// ResponseType : レスポンスの型
type ResponseType struct {
	UUID         string              `json:"uuid"`
	ServiceName  string              `json:"service_name"`
	ResponseData recommend.Recommend `json:"response_data"`
}
