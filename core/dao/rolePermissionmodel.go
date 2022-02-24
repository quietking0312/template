package dao

import (
	"fmt"
	"go.uber.org/zap"
	"server/common/log"
)

// language=sql
const (
	mRolePermissionInsertSql = "insert ignore into m_role_permission_relation(rid, pid) values (:rid, :pid)"
	mRolePermissionSelectSql = "select pid from m_role_permission_relation"
	mRolePermissionDeleteSql = "delete from m_role_permission_relation where rid=?"
)

type RolePermissionModel struct {
}

func (rp RolePermissionModel) Insert(rid int64, pidS []uint32) error {
	var rpTables []MRolePermissionRelationTable
	for _, pid := range pidS {
		rpTable := MRolePermissionRelationTable{
			Rid: rid,
			Pid: pid,
		}
		rpTables = append(rpTables, rpTable)
	}
	if len(rpTables) > 0 {
		_, err := dao.SqlxNameExec(mRolePermissionInsertSql, rpTables)
		if err != nil {
			log.Error("", zap.Error(err))
			return err
		}
	}
	var (
		sqlStr string
		args   []interface{}
		err    error
	)
	if len(pidS) > 0 {
		sqlStr, args, err = dao.In(fmt.Sprintf("%s and pid not in (?)", mRolePermissionDeleteSql), rid, pidS)
		if err != nil {
			log.Error("", zap.Error(err))
			return err
		}
	} else {
		sqlStr, args, err = dao.In(mRolePermissionDeleteSql, rid)
		if err != nil {
			log.Error("", zap.Error(err))
			return err
		}
	}
	if _, err := dao.SqlxExec(sqlStr, args...); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (rp RolePermissionModel) SelectListByRid(rid int64, pidS *[]uint32) error {
	if err := dao.SqlxSelect(pidS, fmt.Sprintf("%s where rid=?", mRolePermissionSelectSql), rid); err != nil {
		log.Error("", zap.Error(err))
		return err
	}
	return nil
}
