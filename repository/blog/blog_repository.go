package blog

import (
	"database/sql"
	"github.com/xvbnm48/blog-go-idol/model/blog"
)

type BlogRepository interface {
	FindAll() ([]blog.Blog, error)
	FindById(id int) (blog.Blog, error)
	Create(req blog.Blog) (blog.Blog, error)
	Update(req blog.Blog) (blog.Blog, error)
	Delete(id int) (blog.Blog, error)
}

type repository struct {
	db *sql.DB
}

func NewBlogRepository(db *sql.DB) BlogRepository {
	return &repository{db}
}

func (r *repository) FindAll() ([]blog.Blog, error) {
	var post []blog.Blog
	rows, err := r.db.Query("SELECT * FROM post")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var p blog.Blog
		err := rows.Scan(&p.Id, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		post = append(post, p)
	}

	return post, nil
}

func (r *repository) FindById(id int) (blog.Blog, error) {
	return blog.Blog{}, nil
}

func (r *repository) Create(req blog.Blog) (blog.Blog, error) {
	return blog.Blog{}, nil
}

func (r *repository) Update(req blog.Blog) (blog.Blog, error) {
	return blog.Blog{}, nil
}

func (r *repository) Delete(id int) (blog.Blog, error) {
	return blog.Blog{}, nil
}
