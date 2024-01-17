package blog

import (
	"github.com/xvbnm48/blog-go-idol/helper"
	md "github.com/xvbnm48/blog-go-idol/model/blog"
	"github.com/xvbnm48/blog-go-idol/repository/blog"
)

type Server struct {
	repo blog.BlogRepository
}

func NewBlogService(repo blog.BlogRepository) *Server {
	return &Server{
		repo: repo,
	}
}

func (s *Server) GetAllPost() ([]md.Blog, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *Server) FindPostById(id helper.GetById) (md.Blog, error) {
	//TODO implement me
	idPost := id.Id
	post, err := s.repo.FindById(idPost)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (s *Server) CreatePost(req md.BlogCreate) (md.BlogCreate, error) {
	//TODO implement me
	var postData md.BlogCreate
	postData.Title = req.Title
	postData.Content = req.Content

	_, err := s.repo.Create(postData)
	if err != nil {
		return md.BlogCreate{}, err
	}

	// panic("implement me")
	return postData, nil
}

func (s *Server) UpdatePost(req md.Blog) (md.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeletePost(id int) (md.Blog, error) {
	//TODO implement me
	panic("implement me")
}
