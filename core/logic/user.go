package logic

import (
	"errors"
	"fmt"
	"server/common/cyptos"
	"server/common/mtime"
	"server/core/dao"
)

type UserLogic struct{}

func (u UserLogic) IsExistUsername() {

}

func (u UserLogic) Login(username string, password string) (string, error) {
	userModel := new(dao.UserModel)
	var userTable dao.MUserTable
	if err := userModel.SelectOneByUsername(username, &userTable); err != nil {
		if err.Error() == dao.ErrSqlNoRows {
			return "", errors.New("username not exists")
		}
		return "", err
	}
	if userTable.Password != cyptos.Get32MD5(password) {
		return "", errors.New("password is err")
	}
	return cyptos.Get32MD5(fmt.Sprintf("%s.%s.%d", username, password, mtime.GetTime())), nil
}
