package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"server/common/cryptos"
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

func (u UserLogic) Login(username string, password string) (string, error) {
	userModel := new(dao.UserModel)
	var userTable dao.MUserTable
	if err := userModel.SelectOneByUsername(username, &userTable); err != nil {
		if err.Error() == dao.ErrSqlNoRows {
			return "", errors.New("username not exists")
		}
		log.Error("", zap.Error(err))
		return "", err
	}
	if userTable.Password != cryptos.Get32MD5(password) {
		return "", errors.New("password is err")
	}
	return cryptos.Get32MD5(fmt.Sprintf("%s.%s.%d", username, password, mtime.GetTime())), nil
}

// GetUserList 该函数返回的用户信息会有 摘要算法后的密码
func (u UserLogic) GetUserList(page, limit int) ([]dao.MUserTable, error) {
	var dest []dao.MUserTable
	userModel := new(dao.UserModel)
	if err := userModel.SelectUserList((page-1)*limit, limit, &dest); err != nil {
		log.Error("", zap.Error(err))
		return nil, err
	}
	return dest, nil
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

func (u UserLogic) AddUser(name, username, password, email string) error {
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
	return userModel.UpdateUserOne(userTable)
}
