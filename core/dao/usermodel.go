package dao

import (
	"fmt"
	"go.uber.org/zap"
	"server/common/log"
	"server/core/utils/define"
)

// language=SQL
const (
	mUserSelectTotalSql     = "select count(*) from m_user where state != ?"
	mUserInsertSql          = "insert into m_user(uid, username, password, name, email, create_time, last_login_time, state) values (:uid, :username, :password, :name, :email, :create_time, :last_login_time, :state)"
	mUserSelectSql          = "select uid, username, password, name, email, create_time, last_login_time, state from m_user"
	mUserUpdateSql          = "update m_user set name=:name, email=:email, state=:state where uid=:uid"
	mUserUpdatePassByUidSql = "update m_user set password=? where uid=?"
	mUserSelectByPidSql     = "select uid, username, password, name, email, create_time, last_login_time, state from m_user where uid in (select uid from m_user_permission_relation where pid=?)"
)

type UserModel struct {
}

func (u UserModel) InsertOne(user MUserTable) error {
	if _, err := dao.SqlxNameExec(mUserInsertSql, user); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u UserModel) SelectOneByUsername(username string, user *MUserTable) error {
	if err := dao.SqlxGet(user, fmt.Sprintf("%s where username=?", mUserSelectSql), username); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u UserModel) SelectUserList(index, limit int, dest *[]MUserTable) error {
	return dao.SqlxSelect(dest, fmt.Sprintf("%s limit %d, %d", mUserSelectSql, index, limit))
}

func (u UserModel) SelectUserTotal(total *int) error {
	return dao.SqlxGet(total, mUserSelectTotalSql, define.UserStateDelete)
}

func (u UserModel) UpdateUserOne(user MUserTable) error {
	if user.Uid == 0 {
		return nil
	}
	if _, err := dao.SqlxNameExec(mUserUpdateSql, user); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u UserModel) UpdateUserPass(uid int64, password string) error {
	if uid == 0 {
		return nil
	}
	if _, err := dao.Exec(mUserUpdatePassByUidSql, password, uid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u UserModel) SelectUserByPid(pid uint32, user *[]MUserTable) error {
	if err := dao.SqlxSelect(user, mUserSelectByPidSql, pid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}
