package recommend

// Recommend : レコメンドレスポンスの型
type Recommend struct {
	Success    bool     `json:"success"`
	Text       string   `json:"text"`
	ImagePaths []string `json:"image_paths"`
}

// Request : リクエスト型
type Request struct {
	// リクエストとして送る型を決める？
}
