package dao

// MUserTable 用户表
type MUserTable struct {
	Uid           int64  `db:"uid" json:"uid"`
	UserName      string `db:"username" json:"username"`
	Password      string `db:"password" json:"password"`
	Name          string `db:"name" json:"name"`
	Email         string `db:"email" json:"email"`
	CreateTime    int64  `db:"create_time" json:"create_time"`
	LastLoginTime int64  `db:"last_login_time" json:"last_login_time"`
	State         int8   `db:"state" json:"state"`
}

// MRoleTable 角色表
type MRoleTable struct {
	Rid      int64  `db:"rid" json:"rid"`
	RoleName string `db:"role_name" json:"role_name"`
}

// MUserPermissionRelationTable 用户权限表
type MUserPermissionRelationTable struct {
	Uid int64  `db:"uid"`
	Pid uint32 `db:"pid"`
}

// MRolePermissionRelationTable 角色权限表
type MRolePermissionRelationTable struct {
	Rid int64  `db:"rid"`
	Pid uint32 `db:"pid"`
}

type MUserRoleRelationTable struct {
	Uid int64 `db:"uid"`
	Rid int64 `db:"rid"`
}

// MLogsTable 日志表
type MLogsTable struct {
	Type       string `db:"type"`
	Model      string `db:"model"`
	Args       string `db:"args"`
	CreateTime int64  `db:"create_time"`
	Uid        int64  `db:"uid"`
}
