package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	UserId    string    `json:"user_id" gorm:"column:user_id;type:varchar(50);PRIMARY_KEY"`
	Username  string    `json:"username" gorm:"column:username;type:varchar(50)"`
	Password  string    `json:"password" gorm:"column:password;type:varchar(50)"`
	Email     string    `json:"email" gorm:"column:email;type:varchar(50)"`
	AvatarUrl string    `json:"avatar_url" gorm:"column:avatar_url;type:varchar(1000)"`
	Type      int       `json:"type" gorm:"column:type;type:tinyint"`
	CreateAt  time.Time `json:"create_at" gorm:"column:create_at;type:timestamp"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at;type:timestamp"`
}

func (User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	user.UserId = uuid.New().String()
	user.CreateAt = time.Now()
	user.UpdateAt = time.Now()
	// TODO: encrypt the password
	return nil
}
