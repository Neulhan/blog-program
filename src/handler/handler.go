package handler

import (
	"net/http"

	"github.com/Neulhan/blog-program/src/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func (h Handler) GetPostList(c *gin.Context) {
	var posts []database.Post
	h.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func (h Handler) GetOnePost(c *gin.Context) {
	pk := c.Param("id")

	var post database.Post
	h.DB.First(&post, pk)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func (h Handler) CreatePost(c *gin.Context) {
	type createPostProps struct {
		Title   string `json:"title" binding:"required"`   // 등록번호
		Content string `json:"content" binding:"required"` // 상호
	}
	var props createPostProps
	if err := c.ShouldBindJSON(&props); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := &database.Post{
		Title:   props.Title,
		Content: props.Content,
	}
	h.DB.Create(post)
	c.JSON(200, gin.H{
		"result": "success",
		"post":   post,
	})
}

func (h Handler) DeleteOnePost(c *gin.Context) {
	pk := c.Param("id")

	var post database.Post
	h.DB.Delete(&post, pk)

	c.JSON(200, gin.H{
		"result": "success",
	})
}
