package log

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"sync"
	"testing"
)

func TestNewLogger(t *testing.T) {
	logcfg := &logConfig{
		MaxSize:    10,
		Compress:   true,
		LogPath:    "",
		MaxBackups: 0,
		MaxAge:     0,
		LogLevel:   "info",
	}

	err := InitLog(
		Path(logcfg.LogPath),
		Level(logcfg.LogLevel),
		Compress(logcfg.Compress),
		MaxSize(logcfg.MaxSize),
		MaxBackups(logcfg.MaxBackups),
		MaxAge(logcfg.MaxAge),
		Format("json"),
	)
	if err != nil {
		fmt.Printf("aaa")
		os.Exit(1)
	}
	//defer func() {
	//	if err := recover(); err != nil {
	//	}
	//	debug.PrintStack()
	//}()
	var g sync.WaitGroup
	go func() {
		g.Add(1)
		defer g.Done()
		for i := 0; i < 30; i++ {
			Debug("Debug", zap.String("Debug", "hello"))
		}

	}()
	go func() {
		g.Add(1)
		defer g.Done()
		for i := 0; i < 30; i++ {
			Info("TestLog", zap.String("test", "wwwwhhhh"))
		}
	}()
	go func() {
		g.Add(1)
		defer g.Done()
		for i := 0; i < 30; i++ {
			Warn("Warn", zap.String("Warn", "wwww"))
		}
	}()
	go func() {
		g.Add(1)
		defer g.Done()
		for i := 0; i < 30; i++ {
			Error("error", zap.String("error", "hello"))
		}
	}()
	//Panic("panic", zap.String("panic", "pppp"))
	//Fatal("fatal", zap.String("fatal", "ssss"))

	err = io.EOF

	err1 := fmt.Errorf("sss")
	Info("111", zap.Error(err1))
	fmt.Println(errors.Unwrap(err1), err1)
	g.Wait()
}
