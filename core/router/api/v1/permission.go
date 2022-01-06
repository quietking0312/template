package v1

import (
	"github.com/gin-gonic/gin"
	"server/core/utils/resp"
)

type (
	getUserListReq struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}
	getUserListRes struct {
		Data  []map[string]interface{}
		Total int
	}
)

func GetUserListApi(c *gin.Context) {
	var respData = getUserListRes{
		Data:  []map[string]interface{}{},
		Total: 0,
	}

	resp.JSON(c, resp.Success, "", respData)
}
