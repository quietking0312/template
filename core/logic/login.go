package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"server/common/cryptos"
	"server/common/log"
	"server/common/mtime"
	"server/core/dao"
	"sync"
)

var LoginLogicObj = LoginLogic{
	data:  sync.Map{},
	index: make(map[int64]string),
}

type LoginLogic struct {
	data  sync.Map         // key 是token, value 是LoginUserInfo
	index map[int64]string // uid -> token
	sync.Mutex
}

const (
	ErrValueType   = "value type error"
	ErrTokenExpire = "token is expire"
)

type LoginUserInfo struct {
	Uid           int64            `json:"uid"`
	UserName      string           `json:"username"`
	Name          string           `json:"name"`
	Email         string           `json:"email"`
	CreateTime    int64            `json:"create_time"`
	LastLoginTime int64            `json:"last_login_time"`
	State         int8             `json:"state"`
	Role          []dao.MRoleTable `json:"role"`
	PermissionIds []uint32         `json:"permission_ids"`
	isPid         bool
	isRole        bool
}

func (l *LoginLogic) login(token string, info LoginUserInfo) {
	l.Lock()
	defer l.Unlock()
	if oldToken, ok := l.index[info.Uid]; ok {
		l.data.Delete(oldToken)
	}
	l.index[info.Uid] = token
	l.data.Store(token, info)
}

func (l *LoginLogic) logout(token string) {
	l.Lock()
	defer l.Unlock()
	if info, ok := l.data.LoadAndDelete(token); ok {
		if infoData, o := info.(LoginUserInfo); o {
			delete(l.index, infoData.Uid)
		}
	}
}

func (l *LoginLogic) expire() {
	l.data.Range(func(token, info interface{}) bool {
		infoData, _ := info.(LoginUserInfo)
		if mtime.GetTime() > infoData.LastLoginTime+24*3600 {
			key, _ := token.(string)
			l.logout(key)
		}
		return true
	})
}

func (l *LoginLogic) RemoveDuplicate(v []uint32) []uint32 {
	toIndex := 0
	p := uint32(0)
	for i, _ := range v {
		c := &v[i]
		if p == *c && i != 0 {
			continue
		}
		if i != toIndex {
			v[toIndex] = *c
		}
		toIndex++
		p = *c
	}
	return v[:toIndex]
}

func (l *LoginLogic) GetLoginUserInfo(token string) (LoginUserInfo, error) {
	info, ok := l.data.Load(token)
	if !ok {
		return LoginUserInfo{}, errors.New(ErrTokenExpire)
	}
	infoData, o := info.(LoginUserInfo)
	if !o {
		log.Error("", zap.Error(errors.New(ErrValueType)))
		return LoginUserInfo{}, errors.New(ErrValueType)
	}
	if len(infoData.PermissionIds) == 0 {
		userPermissionModel := new(dao.UserPermissionModel)
		var pidS []uint32
		if err := userPermissionModel.SelectListByUid(infoData.Uid, &pidS); err != nil {
			return LoginUserInfo{}, err
		}
		infoData.PermissionIds = append(infoData.PermissionIds, pidS...)
	}
	if len(infoData.Role) == 0 {
		var roles []dao.MRoleTable
		roleModel := new(dao.UserRoleModel)
		if err := roleModel.SelectRoleListByUid(infoData.Uid, &roles); err != nil {
			return LoginUserInfo{}, err
		}
		infoData.Role = roles
		var rolePermissionModel = new(dao.RolePermissionModel)

		for _, role := range roles {
			var pidS []uint32
			if err := rolePermissionModel.SelectListByRid(role.Rid, &pidS); err != nil {
				return LoginUserInfo{}, err
			}
			infoData.PermissionIds = append(infoData.PermissionIds, pidS...)
		}
	}
	infoData.PermissionIds = l.RemoveDuplicate(infoData.PermissionIds)
	return infoData, nil
}

func (l *LoginLogic) Login(username string, password string) (string, error) {
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
	info := LoginUserInfo{
		Uid:           userTable.Uid,
		UserName:      userTable.UserName,
		Name:          userTable.Name,
		Email:         userTable.Email,
		CreateTime:    userTable.CreateTime,
		LastLoginTime: mtime.GetTime(),
		State:         userTable.State,
	}
	token := cryptos.Get32MD5(fmt.Sprintf("%s.%s.%d", username, password, info.LastLoginTime))

	l.login(token, info)
	return token, nil
}
