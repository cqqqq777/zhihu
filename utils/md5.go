package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	var has = []byte(str)
	data := md5.Sum(has)
	return fmt.Sprintf("%x", data)
}
