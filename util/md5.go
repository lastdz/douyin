package util

import (
	"crypto/md5"
	"fmt"
)

func GetMd5String(text string) string {
	hash := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", hash)
}
