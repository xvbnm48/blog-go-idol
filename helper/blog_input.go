package helper

type GetById struct {
	Id int `json:"id"`
}

type CreateBlog struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
