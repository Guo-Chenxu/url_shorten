package service

import (
	"context"
	"fmt"
	"time"
	"url_shorten/biz/model/url_shorten"
	"url_shorten/conf"
	"url_shorten/consts"
	"url_shorten/dal/mysql"
	"url_shorten/dal/redis"
	"url_shorten/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
)

const (
	retryTimes = 3
	codeLength = 6
)

func CreateShortURL(ctx context.Context, req *url_shorten.CreateShortUrlReq) (*url_shorten.CreateShortUrlResp, *consts.BizCode) {
	code := getCode(ctx)
	if code == "" {
		return nil, &consts.SystemErr
	}
	if req.ExpireTime <= 0 {
		req.ExpireTime = conf.GetURLShorten().Expire
	}

	shortURL := &mysql.ShortURL{
		OriginURL:  req.OriginURL,
		Code:       code,
		ExpireTime: time.Now().Add(time.Duration(req.ExpireTime) * time.Hour),
	}
	bizCode := mysql.NewShortURLDao().InsertShortURL(ctx, shortURL)
	if bizCode != nil {
		return nil, bizCode
	}
	resp := buildShortUrlResp(shortURL)

	go cacheShortURL(ctx, shortURL)
	go redis.NewBloomFilter().AddElement(ctx, consts.ShortURLBloomFilterName, code)
	return resp, nil
}

func Query(ctx context.Context, req *url_shorten.QueryReq) (*mysql.ShortURL, *consts.BizCode) {
	redisKey := fmt.Sprintf(consts.RedisShortURLKeyPrefix, req.Code)
	jsonData, err := redis.GetVal(ctx, redisKey)
	if err != nil && err != redis.RedisNil {
		hlog.CtxErrorf(ctx, "redis get err: %+v", err)
	}

	var shortURL *mysql.ShortURL
	var bizCode *consts.BizCode
	if jsonData != "" {
		json.Unmarshal([]byte(jsonData), &shortURL)
	}
	if shortURL == nil {
		shortURL, bizCode = mysql.NewShortURLDao().GetShortURLByCode(ctx, req.GetCode(), time.Now())
		if bizCode != nil {
			return nil, bizCode
		}
	}

	if shortURL != nil {
		if shortURL.ExpireTime.Before(time.Now()) {
			go redis.DelKey(ctx, redisKey)
			return nil, &consts.ShortURLExpired
		} else {
			go cacheShortURL(ctx, shortURL)
			return shortURL, nil
		}
	}
	return nil, &consts.ParamError
}

func getCode(ctx context.Context) string {
	now := time.Now()
	for i := 0; i < retryTimes; i++ {
		code := utils.GenerateCode(codeLength)
		if res, err := redis.NewBloomFilter().Exists(ctx, consts.ShortURLBloomFilterName, code); err == nil && !res {
			return code
		}
		if mysql.NewShortURLDao().IsCodeAvailible(ctx, code, now) {
			return code
		}
	}
	return ""
}

func cacheShortURL(ctx context.Context, shortURL *mysql.ShortURL) {
	key := fmt.Sprintf(consts.RedisShortURLKeyPrefix, shortURL.Code)
	data, err := json.Marshal(shortURL)
	if err != nil {
		hlog.CtxErrorf(ctx, "json marshal error: %v", err)
		return
	}
	err = redis.KeySet(ctx, key, string(data), consts.RedisShortURLExpireTime)
	if err != nil {
		hlog.CtxErrorf(ctx, "redis insert error: %v", err)
	}
}

func buildShortUrlResp(shortURL *mysql.ShortURL) *url_shorten.CreateShortUrlResp {
	return &url_shorten.CreateShortUrlResp{
		OriginURL:  shortURL.OriginURL,
		ShortURL:   conf.GetURLShorten().Prefix + "/" + shortURL.Code,
		ExpireTime: shortURL.ExpireTime.Format(consts.TimeFormat),
	}
}
