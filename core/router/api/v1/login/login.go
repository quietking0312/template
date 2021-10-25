package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/cyptos"
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
	fmt.Println(req)

	resp.JSON(c, resp.Success, "", loginRes{Token: cyptos.Get32MD5(req.Password)})
}
