package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"server/common/log"
)

const (
	mRolePermissionInsertSql = "insert ignore into m_role_permission_relation(rid, pid) values (:rid, :pid)"
	mRolePermissionSelectSql = "select pid from m_role_permission_relation"
	mRolePermissionDeleteSql = "delete from m_role_permission_relation where rid=? and pid not in (?)"
)

type RolePermissionModel struct {
}

func (rp RolePermissionModel) Insert(rid int64, pidS []uint32) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	var rpTables []MRolePermissionRelationTable
	for _, pid := range pidS {
		rpTable := MRolePermissionRelationTable{
			Rid: rid,
			Pid: pid,
		}
		rpTables = append(rpTables, rpTable)
	}
	_, err := dao.sqlxDB.NamedExecContext(ctx, mRolePermissionInsertSql, rpTables)
	if err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	sqlStr, args, err := sqlx.In(mRolePermissionDeleteSql, rid, pidS)
	if err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	if _, err := dao.sqlxDB.ExecContext(ctx, sqlStr, args...); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (rp RolePermissionModel) SelectListByRid(rid int64, pidS *[]uint32) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	if err := dao.sqlxDB.SelectContext(ctx, pidS, fmt.Sprintf("%s where rid=?", mRolePermissionSelectSql), rid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}
