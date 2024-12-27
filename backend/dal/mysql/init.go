package mysql

import (
	"context"
	"fmt"
	"url_shorten/conf"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlClient *gorm.DB

func Init(ctx context.Context, mysqlConfig conf.Mysql) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Addr, mysqlConfig.Database)
	mysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("mysql 连接失败, %v", err))
	}
	hlog.CtxInfof(ctx, "mysql init success")
}
