package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/core/config"
	"server/core/logic"
	"server/core/utils/resp"
)

func GetAppInfo(c *gin.Context) {
	data := map[string]interface{}{
		"register": logic.Common.Register(),
		"version":  config.GetConfig().Version,
	}
	resp.JSON(c, resp.Success, "", data)
}

func DownloadDemo(c *gin.Context) {

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "hello.log"))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Writer.Write([]byte("world"))
}
