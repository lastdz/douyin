package db

type FavoriteDb struct {
	Token    string `gorm:"primarykey;auto_increment"`
	Video_id string `gorm:"not null"`
}
