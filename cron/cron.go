package cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"zhihu/dao/mysql"
	"zhihu/dao/redisdao"
	g "zhihu/global"
)

// Cron 定时将redis中的点赞数写进mysql以实现热榜功能
func Cron() {
	c := cron.New()
	_, err := c.AddFunc("@every 1h", RedisToMysqlPost)
	if err != nil {
		panic(err)
	}
	_, err = c.AddFunc("@every 5s", RedisToMysqlComment)
	if err != nil {
		panic(err)
	}
	c.Start()
}

func RedisToMysqlPost() {
	pidList, err := mysql.GetPidList()
	if err != nil {
		g.Logger.Warn(fmt.Sprintf("failed to sync posts stars  err:%v", err))
		return
	}
	for _, v := range pidList {
		stars, err := redisdao.GetPostStars(v)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync pid:%d stars failed err:%v", v, err))
			continue
		}
		err = mysql.SyncPostStars(v, stars)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync pid:%d stars failed err:%v", v, err))
		}
	}
}

func RedisToMysqlComment() {
	cidList, err := mysql.GetCidList()
	if err != nil {
		g.Logger.Warn(fmt.Sprintf("failed to sync comments stars  err:%v", err))
		return
	}
	for _, v := range cidList {
		stars, err := redisdao.GetCommentsStars(v)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync cid:%d stars failed err:%v", v, err))
			continue
		}
		err = mysql.SyncCommentStars(v, stars)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("sync cid:%d stars failed err:%v", v, err))
		}
	}
}
