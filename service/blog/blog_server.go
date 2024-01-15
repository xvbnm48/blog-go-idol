package blog

import (
	md "github.com/xvbnm48/blog-go-idol/model/blog"
	"github.com/xvbnm48/blog-go-idol/repository/blog"
)

type Server struct {
	repo blog.BlogRepository
}

func (s *Server) FindPostById(id int) (md.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) CreatePost(req md.Blog) (md.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdatePost(req md.Blog) (md.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) DeletePost(id int) (md.Blog, error) {
	//TODO implement me
	panic("implement me")
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
