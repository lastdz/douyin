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
