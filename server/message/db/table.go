package db

import (
	"database/sql"
	"time"
)

type MessageDb struct {
	Id         int64          `gorm:"column:id;type: bigint unsigned;primaryKey;auto_increment;not null"`
	Uid        int64          `gorm:"column:uid;type: bigint unsigned;not null;default:0"`
	ToUserId   int64          `gorm:"column:to_user_id;type: bigint unsigned;not null;default:0"`
	Content    sql.NullString `gorm:"column:content;type: varchar(255);not null;default:''"`
	CreateTime time.Time      `gorm:"column:create_time;type: timestamp;default: CURRENT_TIMESTAMP"`
}
