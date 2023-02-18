package db

import (
	"os"
	"testing"
)

func TestGetDb(t *testing.T) {
	addr := os.Getenv("DEV_MYSQL_ADDR")
	user := os.Getenv("DEV_MYSQL_USER")
	passwd := os.Getenv("DEV_MYSQL_PASSWD")

	GetDb(addr, user, passwd)
}
