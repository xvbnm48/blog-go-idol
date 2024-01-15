package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/xvbnm48/blog-go-idol/handler"
	repo "github.com/xvbnm48/blog-go-idol/repository/blog"
	service "github.com/xvbnm48/blog-go-idol/service/blog"
	"log"
)

func main() {
	fmt.Println("hello world")

	// connect to postgres sql
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "fariz", "fariz", "blog_idol")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
		panic(err)
	}
	defer db.Close()

	fmt.Println("connected to database")

	postRepository := repo.NewBlogRepository(db)
	postService := service.NewBlogService(postRepository)
	postHandler := handler.NewBlogHandler(postService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.GET("/", postHandler.GetAllPost)
	api.GET("/post/:id", postHandler.FindPostById)
	port := ":8081"
	router.Run(port)
}
