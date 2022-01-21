package dao

import (
	"fmt"
	"go.uber.org/zap"
	"server/common/log"
	"server/core/utils/define"
)

const (
	mUserSelectTotalSql     = `select count(*) from m_user where state != ?`
	mUserInsertSql          = `insert into m_user(uid, username, password, name, email, create_time, last_login_time, state) values (:uid, :username, :password, :name, :email, :create_time, :last_login_time, :state)`
	mUserSelectSql          = `select uid, username, password, name, email, create_time, last_login_time, state from m_user`
	mUserUpdateSql          = "update m_user set %s where %s"
	mUserUpdatePassByUidSql = "update m_user set password=? where uid=?"
	mUserSelectByPidSql     = "select uid, username, password, name, email, create_time, last_login_time, state from m_user where uid in (select uid from m_user_permission_relation where pid=?)"
)

type UserModel struct {
}

func (u UserModel) InsertOne(user MUserTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if _, err := dao.sqlxDB.NamedExecContext(ctx, mUserInsertSql, user); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u UserModel) SelectOneByUsername(username string, user *MUserTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if err := dao.sqlxDB.GetContext(ctx, user, fmt.Sprintf("%s where username=?", mUserSelectSql), username); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u UserModel) SelectUserList(index, limit int, dest *[]MUserTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	return dao.sqlxDB.SelectContext(ctx, dest, fmt.Sprintf("%s limit %d, %d", mUserSelectSql, index, limit))
}

func (u UserModel) SelectUserTotal(total *int) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	return dao.sqlxDB.GetContext(ctx, total, mUserSelectTotalSql, define.UserStateDelete)
}

func (u UserModel) UpdateUserOne(user MUserTable) error {
	if user.Uid == 0 {
		return nil
	}
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	updateStr := "state=:state"
	if user.Name != "" {
		updateStr = fmt.Sprintf("%s, name=:name", updateStr)
	}
	if user.Email != "" {
		updateStr = fmt.Sprintf("%s, email=:email", updateStr)
	}
	updateStr = fmt.Sprintf(mUserUpdateSql, updateStr, "uid=:uid")
	if _, err := dao.sqlxDB.NamedExecContext(ctx, updateStr, user); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u UserModel) UpdateUserPass(uid int64, password string) error {
	if uid == 0 {
		return nil
	}
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if _, err := dao.sqlxDB.ExecContext(ctx, mUserUpdatePassByUidSql, password, uid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u UserModel) SelectUserByPid(pid uint32, user *[]MUserTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if err := dao.sqlxDB.SelectContext(ctx, user, mUserSelectByPidSql, pid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}
