package msql

import (
	"database/sql"
	"time"
)

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

type Option func(*dbCfg)

func DriveName(drivename string) Option {
	return func(cfg *dbCfg) {
		cfg.DriveName = drivename
	}
}

func DataSourceName(dsn string) Option {
	return func(cfg *dbCfg) {
		cfg.DataSourceName = dsn
	}
}

func MaxIdleConnection(idle int) Option {
	return func(cfg *dbCfg) {
		cfg.MaxIdleConnection = idle
	}
}

func MaxOpenConnection(open int) Option {
	return func(cfg *dbCfg) {
		cfg.MaxOpenConnection = open
	}
}

func MaxQueryTime(query time.Duration) Option {
	return func(cfg *dbCfg) {
		cfg.MaxQueryTime = query
	}
}
