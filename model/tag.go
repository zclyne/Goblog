package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	TagId string `json:"tag_id" gorm:"column:tag_id;type:varchar(50);PRIMARY_KEY"`
	Name  string `json:"name" gorm:"column:name;type:varchar(50)"`
}

func (t *Tag) BeforeCreate(scope *gorm.Scope) error {
	t.TagId = uuid.New().String()
	return nil
}

func (Tag) TableName() string {
	return "tag"
}
