package mlog

import "testing"

func TestInitLog(t *testing.T) {
	InitLog()
	Info("hello world")
}
