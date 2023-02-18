package db

type RelationDb struct {
	Id         int64 `gorm:"primarykey;auto_increment"`
	Userid     int64 `gorm:"not null"`
	Followerid int64 `gorm:"not null"`
}
