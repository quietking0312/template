package v1

import (
	"github.com/gin-gonic/gin"
	"server/core/logic"
	"server/core/utils/define"
	"server/core/utils/reqs"
	"server/core/utils/resp"
	"strconv"
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
		//if config.GetConfig().Server.Mode == "debug" {
		//	resp.JSON(c, resp.Success, "", loginRes{Token: cryptos.Get32MD5(req.Password)})
		//	return
		//}
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", loginRes{
		Token: token,
	})
}

type (
	getUserInfoReq struct {
		Token string `json:"token" form:"token"`
	}
	getUserInfoRes struct {
		Email         string      `json:"email"`
		LastLoginTime int64       `json:"lastLoginTime"`
		Name          string      `json:"name"`
		Roles         interface{} `json:"roles"`
		PermissionIds []uint32    `json:"permissionIds"`
		Uid           string      `json:"uid"`
		UserName      string      `json:"userName"`
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
	var respData = getUserInfoRes{
		Uid:           strconv.FormatInt(info.Uid, 10),
		Name:          info.Name,
		UserName:      info.UserName,
		Email:         info.Email,
		Roles:         info.Role,
		PermissionIds: info.PermissionIds,
	}
	resp.JSON(c, resp.Success, "", respData)
}

type (
	registerReq struct {
		Name     string `json:"name" binding:"required" form:"userName"`
		UserName string `json:"username" binding:"required" form:"username"`
		Password string `json:"password" binding:"required" form:"password"`
		Email    string `json:"email" binding:"required,email" form:"email"`
	}
)

func RegisterApi(c *gin.Context) {
	if !logic.Common.Register() {
		resp.JSON(c, resp.ErrArgs, "404", nil)
		return
	}
	var reqData registerReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	var pids []uint32
	if !logic.Common.AdminExists() {
		pids = append(pids, define.AdminPid)
	}
	userLogic := new(logic.UserLogic)
	if err := userLogic.AddUserAndPid(reqData.Name, reqData.UserName, reqData.Password, reqData.Email, pids); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", map[string]bool{
		"register": logic.Common.Register(),
	})
}
