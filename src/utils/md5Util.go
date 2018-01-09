package utils

import (
	"crypto/md5"
	"fmt"
)

//MD5加密
func Md5(str string) string {
	b := []byte(str)
	hax := md5.Sum(b)
	s := fmt.Sprintf("%x", hax)
	return s
}