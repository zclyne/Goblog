package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Blog struct {
	BlogId    string    `json:"blog_id" gorm:"column:blog_id;type:varchar(50);PRIMARY_KEY"`
	Title     string    `json:"title" gorm:"column:title;type:varchar(50)"`
	Content   string    `json:"content" gorm:"column:content;type:varchar(20000)"`
	ImageUrl  string    `json:"image_url" gorm:"column:image_url;type:varchar(1000)"`
	Published bool      `json:"published" gorm:"column:published;type:boolean"`
	TypeId    string    `json:"type_id" gorm:"column:type_id;type:varchar(50);"`
	CreateAt  time.Time `json:"create_at" gorm:"column:create_at;type:timestamp"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at;type:timestamp"`
}

func (Blog) TableName() string {
	return "blog"
}

func (blog *Blog) BeforeCreate(scope *gorm.Scope) error {
	blog.BlogId = uuid.New().String()
	blog.CreateAt = time.Now()
	blog.UpdateAt = time.Now()
	return nil
}
