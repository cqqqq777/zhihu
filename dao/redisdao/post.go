package redisdao

import (
	"context"
	"time"
	g "zhihu/global"
)

func PostDetail(pid int64) (string, error) {
	cmd := g.Rdb.Get(context.Background(), GetPostDetailKey(pid))
	return cmd.Val(), cmd.Err()
}

func SetPostDetail(pid int64, value interface{}) {
	g.Rdb.SetEx(context.Background(), GetPostDetailKey(pid), value, time.Minute*5)
}
