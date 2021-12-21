package modules

import (
	"errors"
	"fmt"
	"server/common/cyptos"
	"server/common/mtime"
	"server/core/model"
)

type UserModule struct{}

func (u UserModule) IsExistUsername() {

}

func (u UserModule) Login(username string, password string) (string, error) {
	userModel := new(model.UserModel)
	var userTable model.MUserTable
	if err := userModel.SelectOneByUsername(username, &userTable); err != nil {
		return "", err
	}
	if userTable.Password != cyptos.Get32MD5(password) {
		return "", errors.New("password is err")
	}
	return cyptos.Get32MD5(fmt.Sprintf("%s.%s.%d", username, password, mtime.GetTime())), nil
}
