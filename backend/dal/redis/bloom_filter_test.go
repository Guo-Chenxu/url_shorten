package redis

import (
	"context"
	"testing"
	"url_shorten/conf"
	"url_shorten/consts"
)

func TestNewBloomFilter(t *testing.T) {
	conf.TestInit()
	Init(context.Background(), conf.GetRedis())
	
	NewBloomFilter().NewFilter(context.Background(), consts.ShortURLBloomFilterName, consts.ShortURLBloomFilterRate, consts.ShortURLBloomFilterCap)
	t.Logf("bloom filter created, %#v", NewBloomFilter().filters[consts.ShortURLBloomFilterName])
	NewBloomFilter().AddElement(context.Background(), consts.ShortURLBloomFilterName, "test")
	res, err := NewBloomFilter().Exists(context.Background(), consts.ShortURLBloomFilterName, "test")
	if err != nil {
		t.Errorf("exists element error, %v", err)
	}
	t.Logf("element exists: %v", res)
}
