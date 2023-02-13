package user

import (
	"fmt"
	"testing"
)

func TestJwt(t *testing.T) {
	secret := []byte("TestSampleSecret114514")
	js := NewJwtService(secret)

	username := "sakuraMiko"
	tokenStr, err := js.GenerateToken(username)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenStr)

	err = js.ValidateToken(tokenStr, username)
	if err != nil {
		panic(err)
	}

	err = js.ValidateToken(tokenStr, "sakuraMio")
	if err == nil {
		panic("unexpected validation pass")
	}
}
