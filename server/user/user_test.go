package user

import (
	"fmt"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	str := os.Getenv("DEV_MYSQL_ADDR")
	fmt.Println(str)
}

func TestAnother(t *testing.T) {
	str := os.Getenv("Addr")
	fmt.Println(str)
}

func TestStart(t *testing.T) {
	server := UserServer{}
	addr := os.Getenv("DEV_MYSQL_ADDR")
	user := os.Getenv("DEV_MYSQL_USER")
	passwd := os.Getenv("DEV_MYSQL_PASSWD")
	server.Start(addr, user, passwd)
}
