package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"server/core/logic"
	"server/core/utils/reqs"
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
		jsonStr, _ := json.Marshal(row)
		_ = json.Unmarshal(jsonStr, &tempMap)
		delete(tempMap, "password")
		respData.Data = append(respData.Data, tempMap)
	}
	resp.JSON(c, resp.Success, "", respData)
}

type (
	postUserReq struct {
		Name     string `json:"name" binding:"required"`
		UserName string `json:"userName" binding:"required"`
		Password string `json:"password"`
		Email    string `json:"email" binding:"required"`
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
