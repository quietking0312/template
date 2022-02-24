package dao

import (
	"fmt"
	"go.uber.org/zap"
	"server/common/log"
)

// language=sql
const (
	mUserRoleInsertSql      = "insert ignore into m_user_role_relation(uid, rid) values (:uid, :rid)"
	mUserRoleDeleteByUidSql = "delete from m_user_role_relation where uid = ?"
	mRoleSelectByUidSql     = "select rid, role_name from m_role where rid = (select rid from m_user_role_relation where uid=?)"
	mRidSelectByUidSql      = "select rid from m_user_role_relation where uid = ?"
)

type UserRoleModel struct {
}

func (ur UserRoleModel) InsertByUid(uid int64, rids []int64) error {
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
		if _, err := dao.SqlxNameExec(mUserRoleInsertSql, urTables); err != nil {
			log.Error("", zap.Error(err))
			return err
		}
		sqlStr, args, err = dao.In(fmt.Sprintf("%s and rid not in (?)", mUserRoleDeleteByUidSql), uid, rids)
		if err != nil {
			log.Error("", zap.Error(err))
			return err
		}
	} else {
		sqlStr, args, err = dao.In(mUserRoleDeleteByUidSql, uid)
		if err != nil {
			log.Error("", zap.Error(err))
			return err
		}
	}
	if _, err := dao.Exec(sqlStr, args...); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (ur UserRoleModel) SelectRoleListByUid(uid int64, dest *[]MRoleTable) error {
	if err := dao.SqlxSelect(dest, mRoleSelectByUidSql, uid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (ur UserRoleModel) SelectRidByUid(uid int64, rid *[]int64) error {
	if err := dao.SqlxSelect(rid, mRidSelectByUidSql, uid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}
