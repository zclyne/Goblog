package service

import "goblog/model"

func Login(username, password string) model.User {
	user := model.User{}
	db.Where("username = ? AND password = ?", username, password).First(&user)
	return user
}

func Register(user *model.User) {
	db.Create(&user)
}
