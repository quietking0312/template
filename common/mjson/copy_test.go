package mjson

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCopyStatus(t *testing.T) {
	s := struct {
		A string
		B int64
	}{
		A: "hello",
		B: 10,
	}
	var b = struct {
		A string
		B int64
	}{}
	fmt.Println(reflect.ValueOf(s))
	_ = CopyStatus(s, &b)
	fmt.Println(b)
}
