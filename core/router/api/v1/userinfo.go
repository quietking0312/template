package v1

import (
	"github.com/gin-gonic/gin"
	"server/core/logic"
	"server/core/utils/reqs"
	"server/core/utils/resp"
)

type (
	getUserInfoReq struct {
		Token string `json:"token" form:"token"`
	}
)

func GetUserInfo(c *gin.Context) {
	var req getUserInfoReq
	if err := reqs.ShouldBind(c, &req); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
	}
	info, err := logic.LoginLogicObj.GetLoginUserInfo(req.Token)
	if err != nil {
		switch err.Error() {
		case logic.ErrTokenExpire:
			resp.JSON(c, resp.ErrTokenExpire, err.Error(), nil)
			return
		default:
			resp.JSON(c, resp.ErrServer, err.Error(), nil)
			return
		}
	}
	resp.JSON(c, resp.Success, "", info)
}
