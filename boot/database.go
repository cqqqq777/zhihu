package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	g "zhihu/global"
)

func DatabaseInit() {
	MysqlInit()
	RedisInit()
}

func MysqlInit() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		g.Config.Database.Mysql.Username,
		g.Config.Database.Mysql.Password,
		g.Config.Database.Mysql.Addr,
		g.Config.Database.Mysql.Port,
		g.Config.Database.Mysql.DBName,
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	g.Mdb = db
	g.Logger.Info("connect mysql successfully")
}

func RedisInit() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", g.Config.Database.Redis.Addr, g.Config.Database.Redis.Port),
		Password: g.Config.Database.Redis.Password,
		DB:       g.Config.Database.Redis.DB,
		PoolSize: g.Config.Database.Redis.PoolSize,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		g.Logger.Fatal(fmt.Sprintf("connect redis failed err:%v", err))
	}
	g.Rdb = client
	g.Logger.Info("connect redis successfully")
}
