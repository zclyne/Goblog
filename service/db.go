package service

import (
	"fmt"
	"goblog/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	host := config.Configuration.Mysql.Host
	port := config.Configuration.Mysql.Port
	username := config.Configuration.Mysql.Username
	password := config.Configuration.Mysql.Password
	databaseName := config.Configuration.Mysql.DatabaseName
	// must declare err here and use db, err =
	// instead of db, err := , because golang will treat db as a local variable, not the global one
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, databaseName))
	if err != nil {
		log.Fatal(2, "Failed to open mysql database connection", err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	// defer db.Close()
}
