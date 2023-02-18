package db

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
	"testing"
)

func TestInsertUser(t *testing.T) {
	addr := os.Getenv("DEV_MYSQL_ADDR")
	user := os.Getenv("DEV_MYSQL_USER")
	passwd := os.Getenv("DEV_MYSQL_PASSWD")

	GetDb(addr, user, passwd)

	id, err := InsertUser("user1", "12345678")
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			fmt.Println("Duplicate username: ", err)
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("id = %v\n", id)
	}
}

func TestFind(t *testing.T) {
	addr := os.Getenv("DEV_MYSQL_ADDR")
	user := os.Getenv("DEV_MYSQL_USER")
	passwd := os.Getenv("DEV_MYSQL_PASSWD")

	GetDb(addr, user, passwd)

	usr, err := ExistByNameAndPasswd("user1", "12345678")
	if err != nil {
		t.Error("Cannot found user1", err)
	} else {
		fmt.Println("found", usr.Id)
	}
	_, err = ExistByNameAndPasswd("user1", "123456789")
	if err == nil {
		t.Error("Found user1 unexpectedly")
	}
	_, err = ExistByNameAndPasswd("user2", "12345678")
	if err == nil {
		t.Error("Found user2 unexpectedly")
	}
}
