package db

type RelationDb struct {
	Id         int64 `gorm:"primarykey;auto_increment"`
	Userid     int   `gorm:"not null"`
	Followerid int   `gorm:"not null"`
}
