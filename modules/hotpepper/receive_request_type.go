package hotpepper

import "github.com/yuuis/RecommendSystem/models/request"

// ReceiveRequestType : Hotpepper用のPAリクエスト型
type ReceiveRequestType struct {
	request.ReceiveRequestType
	RequestParameter RequestParameter `json:"service_data_value"`
}
