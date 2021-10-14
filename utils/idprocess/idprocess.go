package idprocess

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type AutoId struct {
	MaxId    uint64      // 最大id
	lastTime int64       // 上次生成时间 秒
	lastId   uint64      // 上次生成id
	queue    chan uint64 // 管道
	cancel   context.CancelFunc
}

func New(MaxId uint64) *AutoId {
	ctx, cancel := context.WithCancel(context.Background())
	a := &AutoId{
		MaxId:    MaxId,
		lastId:   1,
		lastTime: 0,
		queue:    make(chan uint64, 4),
		cancel:   cancel,
	}
	go a.process(ctx)
	return a
}

func (ai *AutoId) process(ctx context.Context) {
	for ai.lastId < ai.MaxId {
		select {
		case <-ctx.Done():
			return
		default:
			newTime := time.Now().Unix() // 秒级时间戳
			if ai.lastTime == 0 || time.Unix(ai.lastTime, 0).Format("20060102") != time.Unix(newTime, 0).Format("20060102") {
				ai.lastId = 1
			}
			var newId uint64 = 0
			ai.lastTime = newTime
			position := len(strconv.FormatUint(ai.MaxId, 10))
			if len(strconv.FormatUint(ai.lastId, 10)) < position {
				newId = uint64(ai.lastTime)*ai.power(10, uint64(position)) + ai.lastId
			} else if len(strconv.FormatUint(ai.lastId, 10)) > position {
				fmt.Println("idProcess", zap.Error(fmt.Errorf("id:%d > maxId:%d", ai.lastId, ai.MaxId)))
				continue
			}
			ai.queue <- newId
			ai.lastId++
		}
	}
}

func (ai *AutoId) power(x, n uint64) uint64 {
	var s uint64 = 1
	for n > 0 {
		s *= x
		n--
	}
	return s
}

func (ai *AutoId) Id() uint64 {
	return <-ai.queue
}

func (ai *AutoId) Close() {
	ai.cancel()
	close(ai.queue)
}
