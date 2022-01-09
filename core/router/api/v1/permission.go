package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/mjson"
	"server/core/logic"
	"server/core/utils/define"
	"server/core/utils/reqs"
	"server/core/utils/resp"
	"strconv"
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
	var reqData getUserListReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, "", nil)
		return
	}
	userLogic := new(logic.UserLogic)
	total, err := userLogic.GetUserTotal()
	if err != nil {
		resp.JSON(c, resp.ErrServer, "", nil)
		return
	}
	var respData = getUserListRes{
		Data:  []map[string]interface{}{},
		Total: total,
	}
	if total == 0 {
		resp.JSON(c, resp.Success, "", respData)
		return
	}
	result, err := userLogic.GetUserList(reqData.Page, reqData.Limit)
	if err != nil {
		resp.JSON(c, resp.ErrServer, "", nil)
		return
	}
	for _, row := range result {
		var tempMap map[string]interface{}
		jsonStr, _ := mjson.Marshal(row)
		_ = mjson.Unmarshal(jsonStr, &tempMap)
		tempMap["uid"] = string(tempMap["uid"].(json.Number))
		delete(tempMap, "password")
		respData.Data = append(respData.Data, tempMap)
	}
	resp.JSON(c, resp.Success, "", respData)
}

type (
	postUserReq struct {
		Name     string `json:"name"  binding:"required" form:"name"`
		UserName string `json:"userName" binding:"required" form:"userName"`
		Password string `json:"password" form:"password"`
		Email    string `json:"email" binding:"required,email" form:"email"`
	}
)

func PostUserApi(c *gin.Context) {
	var reqData postUserReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	userLogic := new(logic.UserLogic)
	if err := userLogic.AddUser(reqData.Name, reqData.UserName, reqData.Password, reqData.Email); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}

type (
	putUserReq struct {
		Uid   string `json:"uid" form:"uid" binding:"required"`
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email" binding:"email"`
		State int8   `json:"state" form:"state" binding:"required,oneof=1 2"`
	}
)

func PutUserApi(c *gin.Context) {
	var reqData putUserReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	uid, err := strconv.ParseInt(reqData.Uid, 10, 64)
	if err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	userLogin := new(logic.UserLogic)
	if err := userLogin.UpdateUser(uid, reqData.Name, reqData.Email, reqData.State); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
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

func GetPermissionList(c *gin.Context) {
	resData := map[string][]define.RouteItem{
		"data": define.DefaultPermissionList.GetList(),
	}
	resp.JSON(c, resp.Success, "", resData)
}

type postUserPermissionReq struct {
	Uid           string   `json:"uid" form:"uid" binding:"required"`
	PermissionIds []uint32 `json:"p_ids" form:"p_ids"`
}

func PostUserPermission(c *gin.Context) {
	var reqData postUserPermissionReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	for _, pid := range reqData.PermissionIds {
		if !define.DefaultPermissionList.PidIsExists(pid) {
			resp.JSON(c, resp.ErrArgs, fmt.Sprintf("p_id:%d not is exists", pid), nil)
		}
	}
	uid, err := strconv.ParseInt(reqData.Uid, 10, 64)
	if err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	userLogin := new(logic.UserLogic)
	if err := userLogin.UpdatePermission(uid, reqData.PermissionIds); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}

func PostRolePermission(c *gin.Context) {
	resp.JSON(c, resp.Success, "", nil)
}
