package dao

import (
	"context"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"server/common/msql"
	"server/core/config"
)

const (
	ErrSqlNoRows = "sql: no rows in result set"
)

var dao Dao

type Dao struct {
	sqlxDB *sqlx.DB
	sqlDB  *msql.DB
}

// InitDB 初始化db
func InitDB() error {
	dbCfg := config.GetConfig().Server.DB
	if err := msql.CreateDB(dbCfg.DriveName); err != nil {
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
	dao = Dao{
		sqlDB:  _db,
		sqlxDB: _db.GetSqlxConn(),
	}
	return initTable()
}

func ContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), config.GetConfig().Server.DB.MaxQueryTime)
}

// 初始化表
func initTable() error {
	sqlBytes, err := ioutil.ReadFile(config.GetConfig().Server.SqlPath)
	if err != nil {
		return err
	}
	if string(sqlBytes) != "" {
		ctx, cancel := ContextWithTimeout()
		defer cancel()
		if _, err := dao.sqlxDB.ExecContext(ctx, string(sqlBytes)); err != nil {
			return err
		}
	}
	return nil
}
