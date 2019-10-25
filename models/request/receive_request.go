package request

import "github.com/yuuis/RecommendSystem/models/service"

// ReceiveRequestType : 送られてきたリクエストに付随する値を格納する型
type ReceiveRequestType struct {
	TopicCategory  service.ServiceCategory `json:"topic_category"`
	RequireService bool                    `json:"require_service"`
}
