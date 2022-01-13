package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"server/common/log"
)

const (
	mUserPermissionInsertSql = "insert ignore into m_user_permission_relation(uid, pid) values (:uid, :pid)"
	mUserPermissionSelectSql = "select pid from m_user_permission_relation"
	mUserPermissionDeleteSql = "delete from m_user_permission_relation where uid = ?"
)

type UserPermissionModel struct {
}

func (up UserPermissionModel) Insert(uid int64, pIdS []uint32) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	var upTables []MUserPermissionRelationTable
	for _, pid := range pIdS {
		upTable := MUserPermissionRelationTable{
			Uid: uid,
			Pid: pid,
		}
		upTables = append(upTables, upTable)
	}
	if len(upTables) > 0 {
		if _, err := dao.sqlxDB.NamedExecContext(ctx, mUserPermissionInsertSql, upTables); err != nil {
			return err
		}
	}
	var (
		sqlStr string
		args   []interface{}
		err    error
	)
	if len(pIdS) > 0 {
		sqlStr, args, err = sqlx.In(fmt.Sprintf("%s and pid not in (?)", mUserPermissionDeleteSql), uid, pIdS)
		if err != nil {
			return err
		}
	} else {
		sqlStr, args, err = sqlx.In(fmt.Sprintf("%s", mUserPermissionDeleteSql), uid)
		if err != nil {
			return err
		}
	}

	_, err = dao.sqlxDB.ExecContext(ctx, sqlStr, args...)
	return err
}

func (up UserPermissionModel) SelectListByUid(uid int64, pidS *[]uint32) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if err := dao.sqlxDB.SelectContext(ctx, pidS, fmt.Sprintf("%s where uid=?", mUserPermissionSelectSql), uid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}
