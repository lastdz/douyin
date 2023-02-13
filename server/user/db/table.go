package db

type UserModel struct {
	Id       int64  `gorm:"primarykey;auto_increment"`
	Name     string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
