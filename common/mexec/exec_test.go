package mexec

import (
	"fmt"
	"server/common/mencoding"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	c := NewMExec(Powershell())
	c.Start("ping www.baidu.com -n 50")
}

func TestExecCommandRun(t *testing.T) {

	stdout, stderr := ExecCommand("", CMD, []string{"ls"})
	result, _ := mencoding.Byte2Utf8(stdout)
	fmt.Println(string(result))
	fmt.Println(stderr)
}

func TestMExec_Start(t *testing.T) {
	var tempChan = make(chan int, 10)
	go func() {
		for {
			select {
			case i := <-tempChan:
				fmt.Println("r1:", i)
			}
		}
	}()
	go func() {
		for {
			select {
			case i := <-tempChan:
				fmt.Println("r2:", i)
			}
		}
	}()
	for i := 0; i < 100; i++ {
		tempChan <- i
	}
	time.Sleep(1 * time.Second)
}
