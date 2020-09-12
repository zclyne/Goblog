package model

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Type struct {
	TypeId string `json:"type_id" gorm:"column:type_id;type:varchar(50);PRIMARY_KEY"`
	Name   string `json:"name" gorm:"column:name;type:varchar(50)"`
}

func (t *Type) BeforeCreate(scope *gorm.Scope) error {
	t.TypeId = uuid.New().String()
	return nil
}

func (Type) TableName() string {
	return "type"
}
