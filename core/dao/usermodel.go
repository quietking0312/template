package dao

import (
	"fmt"
	"server/core/utils/define"
)

const (
	mUserSelectTotalSql = `select count(*) from m_user where state != ?`
	mUserInsertSql      = `insert into m_user(uid, username, password, name, email, create_time, last_login_time, state) values (:uid, :username, :password, :name, :email, :create_time, :last_login_time, :state)`
	mUserSelectSql      = `select uid, username, password, name, email, create_time, last_login_time, state from m_user`
	mUserUpdateSql      = "update m_user set %s where %s"
)

type UserModel struct {
}

func (u UserModel) InsertOne(user MUserTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if _, err := dao.sqlxDB.NamedExecContext(ctx, mUserInsertSql, user); err != nil {
		return err
	}
	return nil
}

func (u UserModel) SelectOneByUsername(username string, user *MUserTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	return dao.sqlxDB.GetContext(ctx, user, fmt.Sprintf("%s where username=?", mUserSelectSql), username)
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
	updateStr := fmt.Sprintf("state=%d", user.State)
	if user.Name != "" {
		updateStr = fmt.Sprintf("%s, name='%s'", updateStr, user.Name)
	}
	if user.Password != "" {
		updateStr = fmt.Sprintf("%s, password='%s'", updateStr, user.Password)
	}
	if user.Email != "" {
		updateStr = fmt.Sprintf("%s, email='%s'", updateStr, user.Email)
	}
	updateStr = fmt.Sprintf(mUserUpdateSql, updateStr, fmt.Sprintf("uid=%d", user.Uid))
	fmt.Println(updateStr)
	_, err := dao.sqlxDB.ExecContext(ctx, updateStr)
	return err
}
