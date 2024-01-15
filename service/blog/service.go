package blog

import "github.com/xvbnm48/blog-go-idol/model/blog"

type Service interface {
	GetAllPost() ([]blog.Blog, error)
	FindPostById(id int) (blog.Blog, error)
	CreatePost(req blog.Blog) (blog.Blog, error)
	UpdatePost(req blog.Blog) (blog.Blog, error)
	DeletePost(id int) (blog.Blog, error)
}