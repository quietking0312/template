package reqs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/common/log"
	"server/core/config"
)

func ShouldBind(c *gin.Context, req interface{}) error {
	if err := c.ShouldBind(req); err != nil {
		return err
	}
	if config.GetConfig().Server.Mode == "debug" {
		log.Info(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL), zap.Any("request", req))
	}
	return nil
}
