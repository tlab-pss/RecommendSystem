package service

// ServiceCategory : サービスのカテゴリを保持する（個別に持たない方法ないかなぁ……）
type ServiceCategory int

// Todo : ていうかpackage読み込めるなら、PSScore的なパッケージ用意してそこに共通モジュール持ちたい

const (
	// Uncategorized : 未分類
	Uncategorized ServiceCategory = iota
	// Commerce : コマース
	Commerce
	// Gourmet : グルメ
	Gourmet
	// Weather : 天気
	Weather
	// Map : マップ
	Map
	// Mail : メール
	Mail
	// Music : ミュージック
	Music
	// Message : メッセージ
	Message
	// Search : 検索
	Search
	// Translation : 翻訳
	Translation
	// News : ニュース
	News
	// PersonalData : 個人情報
	PersonalData
)
