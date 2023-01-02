package g

import (
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"zhihu/model/config"
)

var (
	Config *config.Config
	Mdb    *sqlx.DB
	Rdb    *redis.Client
	Logger *zap.Logger
)
