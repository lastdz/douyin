package db

import (
	"database/sql"
	"fmt"
	"math/rand"
	"testing"
)

func TestInsertMessage(t *testing.T) {
	for i := 0; i <= 100; i++ {
		uid := int64(rand.Intn(10))
		err := InsertMessage(&MessageDb{
			Uid:      uid,
			ToUserId: uid + 1,
			Content:  sql.NullString{String: "test", Valid: true},
		})
		if err != nil {
			panic(err)
		}
	}
}

func TestGetAllMessage(t *testing.T) {
	messagesInfo, err := GetAllMessage(1, 2)
	if err != nil {
		panic(err)
	}
	for _, messages := range messagesInfo {
		fmt.Println(messages)
	}
}
