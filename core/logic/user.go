package logic

import (
	"errors"
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

type UserPidItem struct {
	Uid           int64    `json:"uid"`
	UserName      string   `json:"username"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	CreateTime    int64    `json:"create_time"`
	LastLoginTime int64    `json:"last_login_time"`
	State         int8     `json:"state"`
	Rids          []int64  `json:"rids"`
	PermissionIds []uint32 `json:"permission_ids"`
}

func (u UserLogic) GetUserOneByUsername(username string, dest *dao.MUserTable) error {
	userModel := new(dao.UserModel)
	if err := userModel.SelectOneByUsername(username, dest); err != nil {
		if err.Error() == dao.ErrSqlNoRows {
			return errors.New("username not exists")
		}
		log.Error("", zap.Error(err))
		return err
	}
	return nil
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
	roleUserModel := new(dao.UserRoleModel)
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
		var rids []int64
		if err := roleUserModel.SelectRidByUid(userItem.Uid, &rids); err != nil {
			return nil, err
		}
		userPid.Rids = rids
		userPidList = append(userPidList, userPid)
	}
	return userPidList, nil
}

func (u UserLogic) GetUserAll() ([]UserPidItem, error) {
	// 偷懒 调用分页查询进行查找， 人数超过500是更改方案
	return u.GetUserList(1, 500)
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

func (u UserLogic) AddUserAndPid(name, username, password, email string, pidS []uint32) error {
	userTable, err := u.addUser(name, username, password, email)
	if err != nil {
		return err
	}
	if err := u.UpdatePermission(userTable.Uid, pidS); err != nil {
		return err
	}
	return nil
}

func (u UserLogic) AddUserAndRole(name, username, password, email string, rids []int64) error {
	userTable, err := u.addUser(name, username, password, email)
	if err != nil {
		return err
	}
	if err := u.UpdateRole(userTable.Uid, rids); err != nil {
		return err
	}
	return nil
}

func (u UserLogic) addUser(name, username, password, email string) (dao.MUserTable, error) {
	if password == "" {
		password = define.DefaultPass
	}
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
		return dao.MUserTable{}, err
	}
	return userTable, nil
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

// ResetUserPass 修改密码
// password 为明文， 方法内部会使用 哈希
func (u UserLogic) ResetUserPass(uid int64, password string) error {
	if password == "" {
		password = define.DefaultPass
	}
	userModel := new(dao.UserModel)
	if err := userModel.UpdateUserPass(uid, define.CryptosPass(password)); err != nil {
		return err
	}
	return nil
}

func (u UserLogic) UpdatePermission(uid int64, pidS []uint32) error {
	userPermissionModel := new(dao.UserPermissionModel)
	if err := userPermissionModel.Insert(uid, pidS); err != nil {
		return err
	}
	// 存在admin权限， 修改全局变量
	for _, pid := range pidS {
		if pid == define.AdminPid {
			Common.SetAdminExists(true)
		}
	}
	return nil
}

// GetPidAllByUid 获取用户及用户所拥有角色权限
func (u UserLogic) GetPidAllByUid(uid int64) ([]uint32, error) {
	var pids []uint32
	userPermissionModel := new(dao.UserPermissionModel)
	if err := userPermissionModel.SelectAllByUid(uid, &pids); err != nil {
		return nil, err
	}
	return pids, nil
}

func (u UserLogic) UpdateRole(uid int64, rids []int64) error {
	return u.updateRole(uid, rids)
}

func (u UserLogic) updateRole(uid int64, rids []int64) error {
	userRoleModel := new(dao.UserRoleModel)
	if err := userRoleModel.InsertByUid(uid, rids); err != nil {
		return err
	}
	return nil
}

func (u UserLogic) adminExists() (bool, error) {
	var users []dao.MUserTable
	userModel := new(dao.UserModel)
	if err := userModel.SelectUserByPid(define.AdminPid, &users); err != nil {
		return false, err
	}
	return len(users) > 0, nil
}
