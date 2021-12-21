package model

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"server/common/msql"
	"server/core/config"
)

var model Model

type Model struct {
	sqlxDB *sqlx.DB
	sqlDB  *msql.DB
}

// InitDB 初始化db
func InitDB() error {
	dbCfg := config.GetConfig().Server.DB
	_db, err := msql.NewDb(
		msql.DriveName(dbCfg.DriveName),
		msql.DataSourceName(dbCfg.Dsn),
		msql.MaxQueryTime(dbCfg.MaxQueryTime),
		msql.MaxOpenConnection(dbCfg.MaxConn),
		msql.MaxIdleConnection(dbCfg.MaxIdle))
	if err != nil {
		return err
	}
	model = Model{
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
		if _, err := model.sqlxDB.ExecContext(ctx, string(sqlBytes)); err != nil {
			return err
		}
	}
	return nil
}

type TxOption func(options *sql.TxOptions)

// LevelReadCommitted 读取完成立刻释放共享锁模式
func LevelReadCommitted() TxOption {
	return func(options *sql.TxOptions) {
		options.Isolation = sql.LevelReadCommitted
	}
}

// LevelRepeatableRead 事务完成释放共享锁模式
func LevelRepeatableRead() TxOption {
	return func(options *sql.TxOptions) {
		options.Isolation = sql.LevelRepeatableRead
	}
}

// LevelSerializable 事务序列操作
func LevelSerializable() TxOption {
	return func(options *sql.TxOptions) {
		options.Isolation = sql.LevelSerializable
	}
}

func DefaultTxOptions() *sql.TxOptions {
	return &sql.TxOptions{
		Isolation: sql.LevelDefault,
	}
}
