package utils

import (
	"errors"
	"fmt"
	"time"
)

const (
	workerBits  = 10                              // 工作机器位数
	numberBits  = 12                              // 每个节点每秒可产生的id数量
	workerMax   = ^uint64(0) >> (64 - workerBits) // 最大机器号
	numberMax   = ^uint64(0) >> (64 - numberBits) // 每个节点最大id数量
	timeShift   = workerBits + numberBits         // 时间戳向左的偏移量
	workerShift = numberBits                      // 工作机器号向左偏移量
	epoch       = 1741417863                      // 起始时间戳变量, 减少id浪费. 不可修改, 否则会出现重复id
)

type SnowFlakeWorker struct {
	workerId  uint64
	timestamp uint64
	number    uint64
}

func NewSnowFlakeWorker(workerId uint64) (*SnowFlakeWorker, error) {
	if workerId > workerMax {
		return nil, errors.New("workId 不正确, 工作机器编号不允许大于 " + fmt.Sprint(workerMax))
	}
	return &SnowFlakeWorker{
		workerId: workerId,
	}, nil
}

func (w *SnowFlakeWorker) GenerateId() (uint64, error) {
	now := getNowTime()

	if now == w.timestamp {
		w.number++
		if w.number > numberMax {
			w.number = 0
			for now <= w.timestamp {
				now = getNowTime()
			}
			w.timestamp = now
		}
	} else if now > w.timestamp {
		w.number = 0
		w.timestamp = now
	} else {
		for now < w.timestamp {
			now = getNowTime()
		}
		w.number++
		if w.number > numberMax {
			w.number = 0
			for now <= w.timestamp {
				now = getNowTime()
			}
			w.timestamp = now
		}
	}

	id := (now-epoch)<<timeShift | (w.workerId << workerShift) | w.number
	return id, nil
}

func getNowTime() uint64 {
	return uint64(time.Now().Unix())
}
