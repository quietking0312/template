package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/cyptos"
	"server/core/protocol"
	"server/core/utils/resp"
)

func Login(c *gin.Context) {
	var req protocol.Login
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		resp.JSON(c, resp.Success, "", "")
		return
	}
	fmt.Println(req.String())

	resp.JSON(c, resp.Success, "", cyptos.Get32MD5(req.GetPassword()))
}
