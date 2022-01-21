package v1

import (
	"github.com/gin-gonic/gin"
	"server/core/config"
	"server/core/logic"
	"server/core/utils/resp"
)

func GetAppInfo(c *gin.Context) {
	data := map[string]interface{}{
		"adminExists": logic.Common.AdminExists(),
		"version":     config.GetConfig().Version,
	}
	resp.JSON(c, resp.Success, "", data)
}
