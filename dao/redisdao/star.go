package redisdao

import (
	"context"
	"github.com/go-redis/redis/v9"
	g "zhihu/global"
)

func StarPost(pid, uid int64) error {
	err := g.Rdb.SetBit(context.Background(), GetStarPostKey(pid), uid, 1).Err()
	return err
}

func CancelStarPost(pid, uid int64) error {
	err := g.Rdb.SetBit(context.Background(), GetStarPostKey(pid), uid, 0).Err()
	return err
}

func GetUserStarStatus(pid, uid int64) (int64, error) {
	intCmd := g.Rdb.GetBit(context.Background(), GetStarPostKey(pid), uid)
	return intCmd.Val(), intCmd.Err()
}

func GetPostStars(pid int64) (int64, error) {
	intCmd := g.Rdb.BitCount(context.Background(), GetStarPostKey(pid), &redis.BitCount{})
	return intCmd.Val(), intCmd.Err()
}
