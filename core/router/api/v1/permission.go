package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/mjson"
	"server/core/dao"
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

func GetUserAllApi(c *gin.Context) {
	var userLogic = new(logic.UserLogic)
	result, err := userLogic.GetUserAll()
	if err != nil {
		resp.JSON(c, resp.ErrServer, "", nil)
		return
	}
	var data []map[string]interface{}
	for _, row := range result {
		var item = map[string]interface{}{
			"uid":  strconv.FormatInt(row.Uid, 10),
			"name": row.Name,
		}
		data = append(data, item)
	}
	resp.JSON(c, resp.Success, "", data)
}

type (
	postUserReq struct {
		Name     string  `json:"name"  binding:"required" form:"name"`
		UserName string  `json:"userName" binding:"required" form:"userName"`
		Password string  `json:"password" form:"password"`
		Email    string  `json:"email" binding:"required,email" form:"email"`
		Rids     []int64 `json:"rids" form:"rids"`
	}
)

func PostUserApi(c *gin.Context) {
	var reqData postUserReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	userLogic := new(logic.UserLogic)
	if err := userLogic.AddUserAndRole(reqData.Name, reqData.UserName, reqData.Password, reqData.Email, reqData.Rids); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}

type (
	putUserReq struct {
		Uid   string  `json:"uid" form:"uid" binding:"required"`
		Name  string  `json:"name" form:"name"`
		Email string  `json:"email" form:"email" binding:"email"`
		State int8    `json:"state" form:"state" binding:"required,oneof=1 2"`
		Rids  []int64 `json:"rids" form:"rids"`
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
	if err := userLogin.UpdateRole(uid, reqData.Rids); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}

func DeleteUserApi(c *gin.Context) {

	resp.JSON(c, resp.Success, "", nil)
}

type (
	deletePassReq struct {
		Uid string `json:"uid" form:"uid" binding:"required"`
	}
)

func DeletePassApi(c *gin.Context) {
	var reqData deletePassReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	uid, err := strconv.ParseInt(reqData.Uid, 10, 64)
	if err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	userLogic := new(logic.UserLogic)
	if err := userLogic.ResetUserPass(uid, ""); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}

type (
	putResetPassReq struct {
		Username    string `json:"username" form:"username" binding:"required"`
		Password    string `json:"password" form:"password" binding:"required"`
		OldPassword string `json:"oldPassword" form:"oldPassword" binding:"required"`
	}
)

func PutResetPassApi(c *gin.Context) {
	var reqData putResetPassReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	var dest dao.MUserTable
	userLogic := new(logic.UserLogic)
	if err := userLogic.GetUserOneByUsername(reqData.Username, &dest); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	if dest.Password != define.CryptosPass(reqData.OldPassword) {
		resp.JSON(c, resp.ErrArgs, "password is err", nil)
		return
	}
	if err := userLogic.ResetUserPass(dest.Uid, reqData.Password); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}

type (
	getRoleListReq struct {
		Page  int `json:"page" form:"page" binding:"required"`
		Limit int `json:"limit" form:"limit" binding:"required"`
	}
	getRoleListRes struct {
		Data  []map[string]interface{} `json:"data"`
		Total int                      `json:"total"`
	}
)

func GetRoleListApi(c *gin.Context) {
	var reqData getRoleListReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	roleLogic := new(logic.RoleLogic)
	total, err := roleLogic.GetRoleTotal()
	if err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	var respData = getRoleListRes{
		Data:  []map[string]interface{}{},
		Total: total,
	}
	if total == 0 {
		resp.JSON(c, resp.Success, "", respData)
		return
	}
	result, err := roleLogic.GetRoleList(reqData.Page, reqData.Limit)
	if err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	for _, row := range result {
		var tempMap map[string]interface{}
		jsonStr, _ := mjson.Marshal(row)
		_ = mjson.Unmarshal(jsonStr, &tempMap)
		respData.Data = append(respData.Data, tempMap)
	}
	resp.JSON(c, resp.Success, "", respData)
}

func GetRoleAllApi(c *gin.Context) {
	roleLogic := new(logic.RoleLogic)
	result, err := roleLogic.GetRoleAll()
	if err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	respData := map[string]interface{}{
		"data": result,
	}
	resp.JSON(c, resp.Success, "", respData)
}

type (
	postRoleReq struct {
		RoleName string `form:"role_name" json:"role_name" binding:"required"`
	}
)

func PostRoleApi(c *gin.Context) {
	var reqData postRoleReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	roleLogic := new(logic.RoleLogic)
	if err := roleLogic.AddRole(reqData.RoleName); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}

type (
	putRoleReq struct {
		Rid      int64  `form:"rid" json:"rid" binding:"required"`
		RoleName string `form:"role_name" json:"role_name" binding:"required"`
	}
)

func PutRoleApi(c *gin.Context) {
	var reqData putRoleReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	roleLogic := new(logic.RoleLogic)
	if err := roleLogic.UpdateRole(reqData.Rid, reqData.RoleName); err != nil {
		resp.JSON(c, resp.ErrServer, "", nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}

func DeleteRoleApi(c *gin.Context) {
	resp.JSON(c, resp.Success, "", nil)
}

func GetPermissionListApi(c *gin.Context) {
	resData := map[string][]define.RouteItem{
		"data": define.DefaultPermissionList.GetList(),
	}
	resp.JSON(c, resp.Success, "", resData)
}

type postUserPermissionReq struct {
	Uid           string   `json:"uid" form:"uid" binding:"required"`
	PermissionIds []uint32 `json:"p_ids" form:"p_ids"`
}

func PostUserPermissionApi(c *gin.Context) {
	var reqData postUserPermissionReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	for _, pid := range reqData.PermissionIds {
		if !define.DefaultPermissionList.PidIsExists(pid) {
			resp.JSON(c, resp.ErrArgs, fmt.Sprintf("p_id:%d not is exists", pid), nil)
			return
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

type postRolePermissionReq struct {
	Rid           int64    `form:"rid" json:"rid" binding:"required"`
	PermissionIds []uint32 `form:"p_ids" json:"p_ids"`
}

func PostRolePermissionApi(c *gin.Context) {
	var reqData postRolePermissionReq
	if err := reqs.ShouldBind(c, &reqData); err != nil {
		resp.JSON(c, resp.ErrArgs, err.Error(), nil)
		return
	}
	roleLogic := new(logic.RoleLogic)
	if err := roleLogic.UpdatePermission(reqData.Rid, reqData.PermissionIds); err != nil {
		resp.JSON(c, resp.ErrServer, err.Error(), nil)
		return
	}
	resp.JSON(c, resp.Success, "", nil)
}
