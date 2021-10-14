package idprocess

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	a := 5
	b := a << 2
	c := b >> 2
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", b)
	fmt.Printf("%b\n", c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
