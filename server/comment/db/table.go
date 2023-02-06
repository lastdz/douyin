package db

type CommentDb struct {
	Id         int64 `gorm:"primarykey;auto_increment"`
	Uid        int   `gorm:"not null"`
	Content    string
	CreateDate string
	Vedioid    int
}
