package db

import (
	"gorm.io/driver/mysql"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var mydb *gorm.DB

func GetDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn))
	return db, err
}

func init() {

	mydb, _ = GetDb("root:123456@tcp(127.0.0.1:3306)/douyinfavorite?charset=utf8mb4&parseTime=True&loc=Local")
	mydb.AutoMigrate(&FavoriteDb{})

}
