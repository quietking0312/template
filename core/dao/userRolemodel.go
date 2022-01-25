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
	mRoleSelectByUidSql     = "select rid, role_name from m_role where rid = (select rid from m_user_role_relation where uid=?)"
	mRidSelectByUidSql      = "select rid from m_user_role_relation where uid = ?"
)

type UserRoleModel struct {
}

func (ur UserRoleModel) InsertByUid(uid int64, rids []int64) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	var (
		sqlStr string
		args   []interface{}
		err    error
	)
	if len(rids) > 0 {
		var urTables []MUserRoleRelationTable
		for _, rid := range rids {
			urTable := MUserRoleRelationTable{
				Uid: uid,
				Rid: rid,
			}
			urTables = append(urTables, urTable)
		}
		if _, err := dao.sqlxDB.NamedExecContext(ctx, mUserRoleInsertSql, urTables); err != nil {
			log.Error("", zap.Error(err))
			return err
		}
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

func (ur UserRoleModel) SelectRoleListByUid(uid int64, dest *[]MRoleTable) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if err := dao.sqlxDB.SelectContext(ctx, dest, mRoleSelectByUidSql, uid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (ur UserRoleModel) SelectRidByUid(uid int64, rid *[]int64) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if err := dao.sqlxDB.SelectContext(ctx, rid, mRidSelectByUidSql, uid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}
