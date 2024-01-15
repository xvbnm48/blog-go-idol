package blog

import (
	"github.com/xvbnm48/blog-go-idol/model/blog"
	"log"
)

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
	var post blog.Blog
	post.Id = id
	err := r.db.QueryRow("SELECT * FROM post WHERE id = $1", id).Scan(&post.Id, &post.Title, &post.Content)
	if err != nil {
		log.Print(err)
		return blog.Blog{}, err
	}

	return post, nil

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
