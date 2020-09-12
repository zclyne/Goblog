package service

import (
	"goblog/model"
	"log"
)

func ListCommentsByBlogId(blogId string) []model.Comment {
	var comments []model.Comment
	db.Where("blog_id = ?", blogId).Order("create_at desc").Find(&comments)
	return comments
}

func CreateComment(comment model.Comment) {
	db.Create(&comment)
}

func DeleteCommentById(commentId string) error {
	if err := db.Where("comment_id = ?", commentId).Delete(model.Comment{}).Error; err != nil {
		log.Printf("error occurred while deleting comment, err is: %s\n", err)
		return err
	}
	return nil
}
