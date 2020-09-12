package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	CommentId string `json:"comment_id" gorm:"column:comment_id;type:varchar(50);PRIMARY_KEY"`
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(50)"`
	Email string `json:"email" gorm:"column:email;type:varchar(50)"`
	BlogId string `json:"blog_id" gorm:"column:blog_id;type:varchar(50)"`
	Content string `json:"content" gorm:"column:content;type:varchar(2000)"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url;type:varchar(1000)"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at;type:timestamp"`
}

func (Comment) TableName() string {
	return "comment"
}

func (c *Comment) BeforeCreate(scope *gorm.Scope) error {
	c.CommentId = uuid.New().String()
	c.CreateAt = time.Now()
	return nil
}