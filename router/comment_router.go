package router

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/service"
	"net/http"
)

func ListCommentsByBlogId(c *gin.Context) {
	blogId := c.Query("blog_id")
	comments := service.ListCommentsByBlogId(blogId)
	c.JSON(http.StatusOK, comments)
}

func CreateComment(c *gin.Context) {
	comment := model.Comment{}
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	service.CreateComment(comment)
}

func DeleteComment(c *gin.Context) {
	commentId := c.Param("comment_id")
	service.DeleteCommentById(commentId)
	c.JSON(http.StatusAccepted, nil)
}