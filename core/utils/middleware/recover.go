package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/common/log"
	"server/core/utils/resp"
)

// Recover panic 处理
func Recover() gin.HandlerFunc {
	DefaultErrorWriter := &PanicExceptionRecord{}
	return gin.RecoveryWithWriter(DefaultErrorWriter, func(c *gin.Context, err interface{}) {
		resp.JSON(c, resp.ErrServer, "", fmt.Sprintf("%s", err))
	})
}

type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	errStr := string(b)
	err = errors.New(errStr)
	log.Error("panic", zap.Error(err))
	return len(errStr), err
}
