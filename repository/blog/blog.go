package blog

import (
	"log"

	"github.com/xvbnm48/blog-go-idol/model/blog"
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

func (r *repository) Create(req blog.BlogCreate) (blog.BlogCreate, error) {
	var post blog.BlogCreate
	post.Title = req.Title
	post.Content = req.Content

	// Perbaikan: Hapus koma ekstra setelah $2
	_, err := r.db.Exec("INSERT INTO post (title, content) VALUES ($1, $2)", post.Title, post.Content)

	if err != nil {
		// Perbaikan: Ganti log.Print dengan log.Println untuk mencetak baris baru
		log.Println(err)
		return blog.BlogCreate{}, err
	}

	return post, nil
}

func (r *repository) Update(req blog.Blog) (blog.Blog, error) {
	return blog.Blog{}, nil
}

func (r *repository) Delete(id int) (blog.Blog, error) {
	return blog.Blog{}, nil
}
