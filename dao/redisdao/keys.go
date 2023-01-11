package redisdao

import "strconv"

const (
	AllTopicKey = "topics"
)

func GetVerificationKey(email string) string {
	return email + ":verification"
}

func GetTopicDetailKey(tid int64) string {
	return "topic:" + strconv.FormatInt(tid, 10) + ":detail"
}

func GetPostDetailKey(pid int64) string {
	return "post:" + strconv.FormatInt(pid, 10) + ":detail"
}

func GetStarPostKey(pid int64) string {
	return "post:" + strconv.FormatInt(pid, 10) + ":stars"
}

func GetStarCommentKey(cid int64) string {
	return "comment:" + strconv.FormatInt(cid, 10) + ":stars"
}
