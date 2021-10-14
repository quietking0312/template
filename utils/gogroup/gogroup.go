package gogroup

import (
	"fmt"
	"time"
)

/*
协程组， 用于控制协程数量， 协程启用过多，则阻塞 任务
*/

type Job func(data TaskData)

type TaskData map[interface{}]interface{}

type GoGroup struct {
	num int
	c   chan struct{}
}

func NewGoGroup(n int) *GoGroup {
	if n <= 0 {
		return nil
	}
	gg := &GoGroup{
		num: n,
		c:   make(chan struct{}, n),
	}
	return gg
}

func (g *GoGroup) Run(data TaskData, job Job) error {
	timer := time.NewTimer(30 * time.Second)
	defer timer.Stop()
	select {
	case g.c <- struct{}{}:
	case <-timer.C:
		return fmt.Errorf("timeout")
	}
	go func() {
		job(data)
		<-g.c
	}()
	return nil
}
