package merrors

import (
	"fmt"
	"testing"
)

func f() (r int) {
	defer func() {
		r += 5
	}()
	return 0
}

func TestNew(t *testing.T) {
	fmt.Println(f())
}
