package redisdao

import (
	"context"
	"strconv"
	"time"
	g "zhihu/global"
)

func SetVerification(email string, vCode int32) error {
	return g.Rdb.SetEx(context.Background(), GetVerificationKey(email), vCode, time.Minute*5).Err()
}

func GetVerification(email string) (int64, error) {
	return strconv.ParseInt(g.Rdb.Get(context.Background(), GetVerificationKey(email)).Val(), 10, 32)
}
