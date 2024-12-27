package redis

import (
	"context"
	"errors"
	"sync"
	"url_shorten/utils"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/redis/go-redis/v9"
)

type BloomFilter struct {
	filters map[string]*Filter
	mutex   *sync.RWMutex
}

type Filter struct {
	rate      float64                                // 误判率
	capacity  int64                                  // 预计容量
	size      int64                                  // bit数组长度
	k         int64                                  // hash函数个数
	seeds     []int64                                // 哈希函数种子
	hashFuncs []func(seed int64, value string) int64 // 哈希函数
}

var bloomFiler *BloomFilter

var bloomFilter sync.Once

func NewBloomFilter() *BloomFilter {
	bloomFilter.Do(func() {
		bloomFiler = &BloomFilter{
			filters: make(map[string]*Filter),
			mutex:   &sync.RWMutex{},
		}
	})
	return bloomFiler
}

func (bf *BloomFilter) NewFilter(ctx context.Context, name string, rate float64, capacity int64) {
	bf.mutex.Lock()
	defer bf.mutex.Unlock()

	size, k := utils.CalculateBloomFilterParams(capacity, rate)
	seeds := utils.GeneratePrimes(k)
	hashFuncs := make([]func(seed int64, value string) int64, k)
	for i := range seeds {
		hashFuncs[i] = utils.CreateHash(size)
	}
	filter := &Filter{
		rate:      rate,
		capacity:  capacity,
		size:      size,
		k:         k,
		seeds:     seeds,
		hashFuncs: hashFuncs,
	}
	bf.filters[name] = filter
}

func (bf *BloomFilter) AddElement(ctx context.Context, name string, element string) error {
	filter, ok := bf.filters[name]
	if !ok {
		hlog.Errorf("bloom filter not exist, name:%s", name)
		return errors.New("bloom filter not exist")
	}

	pipline := redisClient.Pipeline()
	for i, hashFunc := range filter.hashFuncs {
		offset := hashFunc(filter.seeds[i], element)
		pipline.SetBit(ctx, name, offset, 1)
	}
	_, err := pipline.Exec(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "add element fail, name:%s, element:%s, err:%v", name, element, err)
		return err
	}
	return nil
}

func (bf *BloomFilter) Exists(ctx context.Context, name string, element string) (bool, error) {
	filter, ok := bf.filters[name]
	if !ok {
		hlog.Errorf("bloom filter not exist, name:%s", name)
		return false, errors.New("bloom filter not exist")
	}

	redisClient.GetBit(ctx, name, 0)
	pipline := redisClient.Pipeline()
	for i, hashFunc := range filter.hashFuncs {
		offset := hashFunc(filter.seeds[i], element)
		pipline.GetBit(ctx, name, offset)
	}
	cmds, err := pipline.Exec(ctx)
	if err != nil {
		hlog.CtxErrorf(ctx, "exists element fail, name:%s, element:%s, err:%v", name, element, err)
		return false, err
	}
	for _, cmd := range cmds {
		if intCmd, ok := cmd.(*redis.IntCmd); ok && intCmd.Err() == nil && intCmd.Val() == 0 {
			return false, nil
		}
	}
	return true, nil
}
