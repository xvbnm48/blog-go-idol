package blog

type Blog struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type BlogCreate struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
