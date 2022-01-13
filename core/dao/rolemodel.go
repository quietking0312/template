package dao

import (
	"fmt"
	"go.uber.org/zap"
	"server/common/log"
)

const (
	mRoleSelectTotalSql = "select count(*) from m_role"
	mRoleInsertSql      = "insert into m_role(rid, role_name)values(:rid, :role_name)"
	mRoleSelectSql      = "select rid, role_name from m_role"
	mRoleUpdateSql      = "update m_role set %s where %s"
)

type RoleModel struct {
}

func (r RoleModel) InsertOne(role MRoleTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if _, err := dao.sqlxDB.NamedExecContext(ctx, mRoleInsertSql, role); err != nil {
		return err
	}
	return nil
}

func (r RoleModel) SelectRoleList(index, limit int, dest *[]MRoleTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if err := dao.sqlxDB.SelectContext(ctx, dest, fmt.Sprintf("%s limit %d, %d", mRoleSelectSql, index, limit)); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (r RoleModel) SelectRoleTotal(total *int) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	return dao.sqlxDB.GetContext(ctx, total, mRoleSelectTotalSql)
}

func (r RoleModel) UpdateRoleOne(role MRoleTable) error {
	if role.Rid == 0 {
		return nil
	}
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	_, err := dao.sqlxDB.NamedQueryContext(ctx, fmt.Sprintf(mRoleUpdateSql, "role_name=:role_name", "rid=:rid"), role)
	return err
}
