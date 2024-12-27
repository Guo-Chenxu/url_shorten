package dal

import (
	"context"
	"url_shorten/conf"
	"url_shorten/dal/mysql"
	"url_shorten/dal/redis"
)

func Init() {
	ctx := context.Background()
	redis.Init(ctx, conf.GetRedis())
	mysql.Init(ctx, conf.GetMysql())
}
