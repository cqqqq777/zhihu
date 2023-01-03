package utils

import "regexp"

const (
	pattern = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
)

// CheckEmail 判断一个字符串是否为合法的邮箱地址
func CheckEmail(email string) bool {
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
