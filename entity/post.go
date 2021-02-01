package repository

// 型定義
type Post struct {
	ID    int64    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
