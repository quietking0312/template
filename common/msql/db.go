package msql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	_db *DB
)

type dbCfg struct {
	DriveName         string
	DataSourceName    string
	MaxIdleConnection int // 连接池中的最大闲置链接
	MaxOpenConnection int // 与数据库建立链接的最大数目
	MaxQueryTime      time.Duration
}

type DB struct {
	DB    *sql.DB
	dbCfg *dbCfg
}

func (_db *DB) GetConn(ctx context.Context) (*sql.Conn, error) {
	conn, err := _db.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	err = conn.PingContext(ctx)
	return conn, err
}

func (_db *DB) GetSqlxConn() *sqlx.DB {
	return sqlx.NewDb(_db.DB, _db.dbCfg.DriveName)
}

func GetConn(ctx context.Context) (*sql.Conn, error) {
	return _db.GetConn(ctx)
}

func GetSqlxConn() *sqlx.DB {
	return _db.GetSqlxConn()
}

func GetDB() *sql.DB {
	return _db.DB
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

func defaultDBOption() *dbCfg {
	return &dbCfg{
		DriveName:         "",
		DataSourceName:    "",
		MaxIdleConnection: 10,
		MaxOpenConnection: 5,
		MaxQueryTime:      3,
	}
}

func NewDb(opts ...Option) (*DB, error) {
	dbCfg := defaultDBOption()
	for _, opt := range opts {
		opt(dbCfg)
	}
	if dbCfg.DriveName == "mysql" {
		_ = createDB(dbCfg.DataSourceName)
	}
	db, err := sql.Open(dbCfg.DriveName, dbCfg.DataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(dbCfg.MaxOpenConnection)
	db.SetMaxIdleConns(dbCfg.MaxIdleConnection)
	db.SetConnMaxIdleTime(dbCfg.MaxQueryTime)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	_db = &DB{
		DB:    db,
		dbCfg: dbCfg,
	}
	return _db, nil
}

// mysql 创建数据库
func createDB(DataSourceName string) (err error) {
	var (
		db *sql.DB
	)
	cfg, _ := mysql.ParseDSN(DataSourceName)
	if cfg.Params["charset"] == "" {
		cfg.Params["charset"] = "utf8mb4"
	}
	source := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err = sql.Open("mysql", source)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(
		fmt.Sprintf("CREATE DATABASE If Not Exists `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin",
			cfg.DBName))
	if err != nil {
		return err
	}
	return nil
}

func (_db *DB) CreateTable(sql string, args ...interface{}) error {
	_, err := _db.DB.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}
