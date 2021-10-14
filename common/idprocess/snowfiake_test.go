package idprocess

import (
	"fmt"
	"testing"
)

func TestNewWorker(t *testing.T) {
	n1, err := NewWorker(1023)
	if err != nil {
		return
	}
	n2, err := NewWorker(1)
	if err != nil {
		return
	}
	for i := 0; i < 10; i++ {
		fmt.Println(n1.GetId())
		fmt.Printf("%b\n", n1.GetId())

	}
	for j := 0; j < 10; j++ {
		fmt.Println(n2.GetId())
	}
	fmt.Println(int64(-1 ^ (-1 << 63)))
	fmt.Println(1e3)
}
