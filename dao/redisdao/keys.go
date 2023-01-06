package redisdao

import "strconv"

const (
	AllTopicKey = "topics"
)

func GetVerificationKey(email string) string {
	return email + ":verification"
}

func GetTopicDetailKey(tid int64) string {
	return "topic:" + strconv.FormatInt(tid, 10)
}
