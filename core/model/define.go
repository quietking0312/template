package model

type UserTable struct {
	Uid           int64
	UserName      string
	Password      string
	Name          string
	Email         string
	CreateTime    int64
	LastLoginTime int64
	State         int8
}
