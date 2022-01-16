package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"server/common/log"
)

const (
	mUserRoleInsertSql      = "insert ignore into m_user_role_relation(uid, rid) values (:uid, :rid)"
	mUserRoleDeleteByUidSql = "delete from m_user_role_relation where uid = ?"
	mRoleSelectByUidSql     = "select m_role.rid, m_role.role_name from m_user_role_relation left join m_role on m_user_role_relation.rid = m_role.rid where uid=?"
)

type UserRoleModel struct {
}

func (ur UserRoleModel) InsertByUid(uid int64, rids []int64) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	var urTables []MUserRoleRelationTable
	for _, rid := range rids {
		urTable := MUserRoleRelationTable{
			Uid: uid,
			Rid: rid,
		}
		urTables = append(urTables, urTable)
	}
	_, err := dao.sqlxDB.NamedExecContext(ctx, mUserRoleInsertSql, urTables)
	if err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	var (
		sqlStr string
		args   []interface{}
	)
	if len(rids) > 0 {
		sqlStr, args, err = sqlx.In(fmt.Sprintf("%s and rid not in (?)", mUserRoleDeleteByUidSql), uid, rids)
		if err != nil {
			log.Error("", zap.Error(err))
			return err
		}
	} else {
		sqlStr, args, err = sqlx.In(mUserRoleDeleteByUidSql, uid)
		if err != nil {
			log.Error("", zap.Error(err))
			return err
		}
	}
	if _, err := dao.sqlxDB.ExecContext(ctx, sqlStr, args...); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (ur UserRoleModel) SelectRoleListByUid(uid int64, dest *[]RoleModel) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if err := dao.sqlxDB.SelectContext(ctx, dest, mRoleSelectByUidSql, uid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}
