package login

import (
	"github.com/gin-gonic/gin"
	"server/core/utils/reqs"
	"server/core/utils/resp"
)

type (
	getUserInfoReq struct {
		Token string `json:"token" form:"token"`
	}
	getUserInfoRes struct {
		PermissionId []int32 `json:"permission_id"`
	}
)

func GetUserInfo(c *gin.Context) {
	var req getUserInfoReq
	if err := reqs.ShouldBind(c, &req); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
	}

	res := getUserInfoRes{
		PermissionId: []int32{100000},
	}

	resp.JSON(c, resp.Success, "", res)
}
