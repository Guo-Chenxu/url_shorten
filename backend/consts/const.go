package consts

import "time"

const (
	RedisShortURLKeyPrefix  = "url_shorten:code:%s"
	RedisShortURLExpireTime = 24 * time.Hour
)

const TimeFormat = "2006-01-02 15:04:05"

const (
	ShortURLBloomFilterName = "url_shorten:code_bloom_filter"
	ShortURLBloomFilterRate = 0.001
	ShortURLBloomFilterCap  = 10000
)
