package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDb(addr string, user string, passwd string) {
	var err error
	url := fmt.Sprintf("%s:%s@tcp(%s)/douyin_user?charset=utf8mb4&parseTime=True&loc=Local", user, passwd, addr)
	db, err = gorm.Open(mysql.Open(url))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(UserModel{})
}
