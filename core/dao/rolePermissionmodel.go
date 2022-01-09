package dao

import "fmt"

const (
	mRolePermissionInsertSql = "insert ignore into m_role_permission_relation(rid, pid) values (:rid, :pid)"
	mRolePermissionSelectSql = "select pid from m_role_permission_relation"
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
	return err
}

func (rp RolePermissionModel) SelectListByRid(rid int64, pidS *[]int32) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	return dao.sqlxDB.SelectContext(ctx, pidS, fmt.Sprintf("%s where rid=?", mRolePermissionSelectSql), rid)
}
