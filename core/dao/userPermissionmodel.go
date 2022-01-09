package dao

import "fmt"

const (
	mUserPermissionInsertSql = "insert ignore into m_user_permission_relation(uid, pid) values (:uid, :pid)"
	mUserPermissionSelectSql = "select pid from m_user_permission_relation"
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
	_, err := dao.sqlxDB.NamedExecContext(ctx, mUserPermissionInsertSql, upTables)
	return err
}

func (up UserPermissionModel) SelectListByUid(uid int64, pidS *[]uint32) error {
	ctx, cancel := ContextWithTimeout()
	defer cancel()
	return dao.sqlxDB.SelectContext(ctx, pidS, fmt.Sprintf("%s where uid=?", mUserPermissionSelectSql), uid)
}
