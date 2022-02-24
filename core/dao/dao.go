package dao

import (
	"io/ioutil"
	"server/common/msql"
	"server/core/config"
)

const (
	ErrSqlNoRows = "sql: no rows in result set"
)

var dao *msql.DB

// InitDB 初始化db
func InitDB() error {
	dbCfg := config.GetConfig().Server.DB
	if err := msql.CreateDB(dbCfg.Dsn); err != nil {
		return err
	}
	_db, err := msql.NewDb(
		msql.DriveName(dbCfg.DriveName),
		msql.DataSourceName(dbCfg.Dsn),
		msql.MaxQueryTime(dbCfg.MaxQueryTime),
		msql.MaxOpenConnection(dbCfg.MaxConn),
		msql.MaxIdleConnection(dbCfg.MaxIdle))
	if err != nil {
		return err
	}
	dao = _db
	return initTable()
}

// 初始化表
func initTable() error {
	sqlBytes, err := ioutil.ReadFile(config.GetConfig().Server.SqlPath)
	if err != nil {
		return err
	}
	if string(sqlBytes) != "" {
		if _, err := dao.SqlxExec(string(sqlBytes)); err != nil {
			return err
		}
	}
	return nil
}
