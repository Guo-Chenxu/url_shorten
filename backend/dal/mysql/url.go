package mysql

import (
	"context"
	"sync"
	"time"
	"url_shorten/consts"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
)

type ShortURL struct {
	ID         uint      `gorm:"primary_key;auto_increment" json:"id"`
	OriginURL  string    `gorm:"type:text;not null" json:"origin_url"`
	Code       string    `gorm:"type:varchar(255);not null" json:"code"`
	ExpireTime time.Time `gorm:"type:datetime;not null" json:"expire_time"`
	CreateTime time.Time `gorm:"type:datetime;not null" json:"create_time"`
}

type ShortURLDao struct {
}

var shortURLDao *ShortURLDao

var shortURLDaoOnce sync.Once

func NewShortURLDao() *ShortURLDao {
	shortURLDaoOnce.Do(func() {
		shortURLDao = &ShortURLDao{}
	})
	return shortURLDao
}

const ShortURLTableName = "short_url"

func (d *ShortURLDao) InsertShortURL(ctx context.Context, shortURL *ShortURL) *consts.BizCode {
	shortURL.CreateTime = time.Now()
	res := mysqlClient.Table(ShortURLTableName).Create(shortURL)
	if res.Error != nil {
		hlog.CtxErrorf(ctx, "write db err, table: %s, %v", ShortURLTableName, res.Error)
		return &consts.BizCode{Code: consts.WriteDbError.Code, Msg: consts.WriteDbError.Msg}
	}
	return nil
}

func (d *ShortURLDao) GetShortURLByCode(ctx context.Context, code string, expireTime time.Time) (*ShortURL, *consts.BizCode) {
	var shortURL ShortURL
	err := mysqlClient.Debug().Table(ShortURLTableName).Where("code = ?", code).
		Where("expire_time >= ?", expireTime).First(&shortURL).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		hlog.CtxErrorf(ctx, "query db err, table: %s, %v", ShortURLTableName, err)
		return nil, &consts.BizCode{Code: consts.ReadDbError.Code, Msg: consts.ReadDbError.Msg}
	}
	return &shortURL, nil
}

func (d *ShortURLDao) IsCodeAvailible(ctx context.Context, code string, expireTime time.Time) bool {
	var cnt int64
	err := mysqlClient.Debug().Table(ShortURLTableName).Where("code = ?", code).
		Where("expire_time >= ?", expireTime).Count(&cnt).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "query db err, table: %s, %v", ShortURLTableName, err)
	}
	return cnt == 0
}
