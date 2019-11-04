package recommend

// Recommend : レコメンド内容の結果
type Recommend struct {
	Success    bool     `json:"success"`
	Text       string   `json:"text"`
	ImagePaths []string `json:"image_paths"`
}
