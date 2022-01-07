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
		Data  []map[string]interface{} `json:"data"`
		Total int                      `json:"total"`
	}
)

func GetUserListApi(c *gin.Context) {
	var respData = getUserListRes{
		Data:  []map[string]interface{}{},
		Total: 0,
	}

	resp.JSON(c, resp.Success, "", respData)
}

func PostUserApi(c *gin.Context) {

	resp.JSON(c, resp.Success, "", nil)
}

func PutUserApi(c *gin.Context) {

	resp.JSON(c, resp.Success, "", nil)
}

func DeleteUserApi(c *gin.Context) {

	resp.JSON(c, resp.Success, "", nil)
}

func GetRoleListApi(c *gin.Context) {
	resp.JSON(c, resp.Success, "", nil)
}

func PostRoleApi(c *gin.Context) {
	resp.JSON(c, resp.Success, "", nil)
}

func PutRoleApi(c *gin.Context) {
	resp.JSON(c, resp.Success, "", nil)
}

func DeleteRoleApi(c *gin.Context) {
	resp.JSON(c, resp.Success, "", nil)
}
