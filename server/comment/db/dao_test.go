package db

import "testing"

func TestInsertComment(t *testing.T) {
	var com CommentDb
	for i := 0; i <= 100; i++ {
		com = CommentDb{ //
			0, 1, "11", "111", 1,
		}
		InsertComment(&com)
	}
}

func TestDeleteComment(t *testing.T) {
	for i := 1; i <= 103; i++ {
		DeleteComment(int64(i))
	}
}
