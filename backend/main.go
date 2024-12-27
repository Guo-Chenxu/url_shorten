// Code generated by hertz generator.

package main

import (
	"context"
	handler "url_shorten/biz/handler"
	"url_shorten/conf"
	"url_shorten/consts"
	"url_shorten/dal"
	"url_shorten/dal/redis"
	"url_shorten/logger"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/cors"
)

func main() {
	conf.InitConfig()
	logger.Init(conf.GetLogger())
	dal.Init()
	redis.NewBloomFilter().NewFilter(context.Background(), consts.ShortURLBloomFilterName, consts.ShortURLBloomFilterRate, consts.ShortURLBloomFilterCap)

	h := server.Default(server.WithHostPorts(conf.GetConfig().Server.Port))
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length", "Accept", "X-Requested-With"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		}}),
		recovery.Recovery(recovery.WithRecoveryHandler(RecoveryHandler)))
	register(h)
	h.Spin()
}

func RecoveryHandler(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
	defer func() {
		if r := recover(); r != nil {
			hlog.CtxErrorf(c, "[Recovery] panic recovered: %v", r)
		}
	}()
	hlog.CtxErrorf(c, "[Recovery] err=%v\nstack=%s", err, stack)
	hlog.CtxErrorf(c, "Client: %s", ctx.Request.Header.UserAgent())
	base := handler.BaseHandler{}
	base.ErrorResponse(c, ctx, &consts.SystemErr, nil)
	ctx.Abort()
}