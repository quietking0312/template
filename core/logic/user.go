package logic

import (
	"go.uber.org/zap"
	"server/common/idprocess"
	"server/common/log"
	"server/common/mtime"
	"server/core/dao"
	"server/core/utils/define"
)

type UserLogic struct{}

// 用户id 生成器
var uidProcess, _ = idprocess.NewWorker(0)

func (u UserLogic) IsExistUsername() {
}

type UserPidItem struct {
	Uid           int64         `json:"uid"`
	UserName      string        `json:"username"`
	Name          string        `json:"name"`
	Email         string        `json:"email"`
	CreateTime    int64         `json:"create_time"`
	LastLoginTime int64         `json:"last_login_time"`
	State         int8          `json:"state"`
	Role          []RolePidItem `json:"role"`
	PermissionIds []uint32      `json:"permission_ids"`
}

// GetUserList 该函数返回的用户信息会有 摘要算法后的密码
func (u UserLogic) GetUserList(page, limit int) ([]UserPidItem, error) {
	var dest []dao.MUserTable
	userModel := new(dao.UserModel)
	if err := userModel.SelectUserList((page-1)*limit, limit, &dest); err != nil {
		log.Error("", zap.Error(err))
		return nil, err
	}
	var userPidList []UserPidItem
	userPermissionModel := new(dao.UserPermissionModel)
	for _, userItem := range dest {
		var userPid = UserPidItem{
			Uid:           userItem.Uid,
			UserName:      userItem.UserName,
			Name:          userItem.Name,
			Email:         userItem.Email,
			CreateTime:    userItem.CreateTime,
			LastLoginTime: userItem.LastLoginTime,
			State:         userItem.State,
		}
		var pIds []uint32
		if err := userPermissionModel.SelectListByUid(userItem.Uid, &pIds); err != nil {
			return nil, err
		}
		userPid.PermissionIds = pIds
		userPidList = append(userPidList, userPid)
	}
	return userPidList, nil
}

func (u UserLogic) GetUserAll() ([]UserPidItem, error) {
	return nil, nil
}

func (u UserLogic) GetUserTotal() (int, error) {
	var total int
	userModel := new(dao.UserModel)
	if err := userModel.SelectUserTotal(&total); err != nil {
		log.Error("", zap.Error(err))
		return 0, err
	}
	return total, nil
}

func (u UserLogic) AddUser(name, username, password, email string, rids []int64) error {
	userModel := new(dao.UserModel)
	var userTable = dao.MUserTable{
		Uid:           uidProcess.GetId(),
		Name:          name,
		UserName:      username,
		Password:      define.CryptosPass(password),
		Email:         email,
		CreateTime:    mtime.GetTime(),
		LastLoginTime: 0,
		State:         define.UserStateOn,
	}
	if err := userModel.InsertOne(userTable); err != nil {
		return err
	}
	if err := u.UpdateRole(userTable.Uid, rids); err != nil {
		return err
	}
	return nil
}

func (u UserLogic) UpdateUser(uid int64, name, email string, state int8) error {
	userModel := new(dao.UserModel)
	var userTable = dao.MUserTable{
		Uid:   uid,
		Name:  name,
		Email: email,
		State: state,
	}
	if err := userModel.UpdateUserOne(userTable); err != nil {
		return err
	}
	return nil
}

func (u UserLogic) UpdatePermission(uid int64, pidS []uint32) error {
	userPermissionModel := new(dao.UserPermissionModel)
	return userPermissionModel.Insert(uid, pidS)
}

// GetPidAllByUid 获取用户及用户所拥有角色权限
func (u UserLogic) GetPidAllByUid(uid int64) ([]uint32, error) {
	return nil, nil
}

func (u UserLogic) UpdateRole(uid int64, rids []int64) error {
	userRoleModel := new(dao.UserRoleModel)
	if err := userRoleModel.InsertByUid(uid, rids); err != nil {
		return err
	}
	return nil
}
