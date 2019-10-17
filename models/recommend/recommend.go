package recommend

import "main/models/service"

// Recommend : レコメンドレスポンスの型
type Recommend struct {
	Success    bool     `json:"success"`
	Text       string   `json:"text"`
	ImagePaths []string `json:"image_paths"`
}

// ReceiveRequestType : 送られてきたリクエストに付随する値を格納する型
type ReceiveRequestType struct {
	TopicCategory    service.ServiceCategory `json:"topic_category"`
	RequireService   bool                    `json:"require_service"`
	ServiceDataValue interface{}             `json:"service_data_value"`
}
