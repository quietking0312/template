package gogroup

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewGoGroup(t *testing.T) {
	var wg = sync.WaitGroup{}
	g := NewGoGroup(2)
	for i := 0; i < 30; i++ {
		wg.Add(1)
		goFunc := func(data TaskData) {
			fmt.Printf("go func: %d\n", data["i"])
			time.Sleep(33 * time.Second)
			wg.Done()
		}
		err := g.Run(TaskData{"i": i}, goFunc)
		if err != nil {
			fmt.Println(err)
			wg.Done()
		}
	}
	wg.Wait()
}
