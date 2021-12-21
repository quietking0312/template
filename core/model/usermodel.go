package model

import "fmt"

const (
	mUserInsertSql = `insert into m_user(uid, username, password, name, email, create_time, last_login_time, state) values (:uid, :username, :password, :name, :email, :create_time, :last_login_time, :state)`
	mUserSelectSql = `select uid, username, password, name, email, create_time, last_login_time, state from m_user`
)

type UserModel struct {
}

func (u UserModel) InsertOne(user MUserTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if _, err := model.sqlxDB.NamedExecContext(ctx, mUserInsertSql, user); err != nil {
		return err
	}
	return nil
}

func (u UserModel) SelectOneByUsername(username string, user *MUserTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	return model.sqlxDB.GetContext(ctx, user, fmt.Sprintf("%s where username=?", mUserSelectSql), username)
}
