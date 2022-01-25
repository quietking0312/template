package logic

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"server/common/cryptos"
	"server/common/log"
	"server/common/mtime"
	"server/core/dao"
	"server/core/utils/define"
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
		if !l.isExpire(infoData.LastLoginTime) {
			key, _ := token.(string)
			l.logout(key)
		}
		return true
	})
}

// 判断是否过期， 过期返回 false
func (l *LoginLogic) isExpire(t int64) bool {
	return mtime.GetTime() <= t+24*3600
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
	if !l.isExpire(infoData.LastLoginTime) {
		return LoginUserInfo{}, errors.New(ErrTokenExpire)
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
	userLogic := new(UserLogic)
	var userTable dao.MUserTable
	if err := userLogic.GetUserOneByUsername(username, &userTable); err != nil {
		return "", err
	}
	if userTable.Password != cryptos.Get32MD5(password) {
		return "", errors.New("password is err")
	}
	if userTable.State != define.UserStateOn {
		return "", errors.New("账号已禁用")
	}
	if password == define.DefaultPass {
		return "", errors.New("请先修改密码")
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
