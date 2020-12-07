package service

import (
	"crypto/md5"
	"fmt"
)

func MD5INIT(password string) string {
	md5string := fmt.Sprintf("%x",md5.Sum([]byte(password)))
	return md5string
}
