package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/blog-go-idol/helper"
	md "github.com/xvbnm48/blog-go-idol/model/blog"
	"github.com/xvbnm48/blog-go-idol/service/blog"
)

type BlogHandler struct {
	blogervice blog.Service
}

func NewBlogHandler(blogService blog.Service) *BlogHandler {
	return &BlogHandler{
		blogervice: blogService,
	}
}
func (h *BlogHandler) FindPostById(c *gin.Context) {
	var input helper.GetById
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	input.Id = idInt
	//err := c.ShouldBindJSON(&input)
	//if err != nil {
	//	c.JSON(400, gin.H{
	//		"message": "error bind json",
	//	})
	//	return
	//}
	Post, err := h.blogervice.FindPostById(input)
	if err != nil {
		response := helper.ApiResponse("data not found", "error", 400, nil)
		c.JSON(400, response)
		return
	}

	response := helper.ApiResponse("success get post", "success", 200, Post)
	c.JSON(200, response)
}
func (h *BlogHandler) GetAllPost(c *gin.Context) {
	//var input []md.Blog
	//err := c.ShouldBindJSON(&input)
	//if err != nil {
	//	c.JSON(400, gin.H{
	//		"message": "error",
	//	})
	//	return
	//}

	Posts, err := h.blogervice.GetAllPost()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error",
		})
		return
	}
	response := helper.ApiResponse("success get all post", "success", 200, Posts)
	c.JSON(200, response)
}

func (h *BlogHandler) CreateNewPost(c *gin.Context) {
	var input md.BlogCreate
	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "error input, please check again",
		})
		return
	}

	_, err = h.blogervice.CreatePost(input)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "error create new post",
		})
		return
	}

	response := helper.ApiResponse("success create new post", "success", 200, nil)
	c.JSON(200, response)
}
