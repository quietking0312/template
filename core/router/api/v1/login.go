package v1

import (
	"github.com/gin-gonic/gin"
	"server/common/cryptos"
	"server/core/config"
	"server/core/logic"
	"server/core/utils/reqs"
	"server/core/utils/resp"
)

type (
	loginReq struct {
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	loginRes struct {
		Token string `json:"token" form:"token"`
	}
)

func Login(c *gin.Context) {
	var req loginReq
	if err := reqs.ShouldBind(c, &req); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), "")
		return
	}
	token, err := logic.LoginLogicObj.Login(req.Username, req.Password)
	if err != nil {
		if config.GetConfig().Server.Mode == "debug" {
			resp.JSON(c, resp.Success, "", loginRes{Token: cryptos.Get32MD5(req.Password)})
			return
		}
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", loginRes{
		Token: token,
	})
}
