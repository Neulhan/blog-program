package main

import (
	"github.com/Neulhan/blog-program/src/database"
	"github.com/Neulhan/blog-program/src/handler"
	"github.com/Neulhan/blog-program/src/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.DBLoader()
	h := handler.NewHandler(db)

	router := gin.Default()
	router.Use(middleware.CORS("*"))

	router.GET("/posts/", h.GetPostList)

	router.POST("/post/create", h.CreatePost)
	router.GET("/post/:id", h.GetOnePost)
	router.DELETE("/post/:id", h.DeleteOnePost)

	router.Run()
}
