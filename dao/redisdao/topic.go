package redisdao

import (
	"context"
	"time"
	g "zhihu/global"
)

func SetAllTopic(value interface{}) {
	g.Rdb.SetEx(context.Background(), AllTopicKey, value, 5*time.Minute)
}

func GetAllTopic() (string, error) {
	cmd := g.Rdb.Get(context.Background(), AllTopicKey)
	return cmd.Val(), cmd.Err()
}

func SetTopicDetail(tid int64, value interface{}) {
	g.Rdb.SetEx(context.Background(), GetTopicDetailKey(tid), value, 5*time.Minute)
}

func GetTopicDetail(tid int64) (string, error) {
	cmd := g.Rdb.Get(context.Background(), GetTopicDetailKey(tid))
	return cmd.Val(), cmd.Err()
}
