package dao

import (
	"fmt"
	"go.uber.org/zap"
	"server/common/log"
)

// language=sql
const (
	mRoleSelectTotalSql = "select count(*) from m_role"
	mRoleInsertSql      = "insert into m_role(rid, role_name)values(:rid, :role_name)"
	mRoleSelectSql      = "select rid, role_name from m_role"
	mRoleUpdateSql      = "update m_role set role_name=:role_name where rid=:rid"
)

type RoleModel struct {
}

func (r RoleModel) InsertOne(role MRoleTable) error {
	if _, err := dao.sqlDB.SqlxNameExec(mRoleInsertSql, role); err != nil {
		return err
	}
	return nil
}

func (r RoleModel) SelectRoleList(index, limit int, dest *[]MRoleTable) error {
	if err := dao.sqlDB.SqlxSelect(dest, fmt.Sprintf("%s limit %d, %d", mRoleSelectSql, index, limit)); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (r RoleModel) SelectRoleTotal(total *int) error {
	return dao.sqlDB.SqlxGet(total, mRoleSelectTotalSql)
}

func (r RoleModel) UpdateRoleOne(role MRoleTable) error {
	if role.Rid == 0 {
		return nil
	}
	_, err := dao.sqlDB.SqlxNameExec(mRoleUpdateSql, role)
	if err != nil {
		log.Error("", zap.Error(err))
	}
	return err
}
