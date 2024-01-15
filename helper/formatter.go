package helper

import "github.com/xvbnm48/blog-go-idol/model/blog"

type postFormatter struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func FormatPost(post blog.Blog) postFormatter {
	formatter := postFormatter{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
	}
	return formatter
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func ApiResponse(message, status string, code int, data any) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
