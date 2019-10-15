package recommend

// Recommend : レコメンドレスポンスの型
type Recommend struct {
	Success bool `json:"success"`
	Text string `json:"text"`
	ImagePaths []string `json:"image_paths"`
}

