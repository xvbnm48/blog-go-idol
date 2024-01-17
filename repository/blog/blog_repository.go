package blog

import (
	"database/sql"

	"github.com/xvbnm48/blog-go-idol/model/blog"
)

type BlogRepository interface {
	FindAll() ([]blog.Blog, error)
	FindById(id int) (blog.Blog, error)
	Create(req blog.BlogCreate) (blog.BlogCreate, error)
	Update(req blog.Blog) (blog.Blog, error)
	Delete(id int) (blog.Blog, error)
}

type repository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) BlogRepository {
	return &repository{db}
}
