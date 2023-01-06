package redisdao

import (
	"context"
	"time"
	g "zhihu/global"
)

func PostDetail(pid int64) (string, error) {
	cmd := g.Rdb.Get(context.Background(), GetTopicDetailKey(pid))
	return cmd.Val(), cmd.Err()
}

func SetPostDetail(pid int64, value interface{}) {
	g.Rdb.SetEx(context.Background(), GetTopicDetailKey(pid), value, time.Minute*5)
}
