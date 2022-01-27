package msql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	_db *DB
)

type DB struct {
	DB     *sql.DB
	SqlxDB *sqlx.DB
	dbCfg  *dbCfg
}

func (_db *DB) GetConn(ctx context.Context) (*sql.Conn, error) {
	conn, err := _db.DB.Conn(ctx)
	if err != nil {
		return nil, err
	}
	err = conn.PingContext(ctx)
	return conn, err
}

func GetConn(ctx context.Context) (*sql.Conn, error) {
	return _db.GetConn(ctx)
}

func GetDB() *sql.DB {
	return _db.DB
}

func NewDb(opts ...Option) (*DB, error) {
	dbCfg := defaultDBOption()
	for _, opt := range opts {
		opt(dbCfg)
	}
	db, err := sql.Open(dbCfg.DriveName, dbCfg.DataSourceName)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(dbCfg.MaxOpenConnection)
	db.SetMaxIdleConns(dbCfg.MaxIdleConnection)
	db.SetConnMaxIdleTime(dbCfg.MaxQueryTime)
	for i := 0; i < dbCfg.MaxIdleConnection; i++ {
		if err = db.Ping(); err != nil {
			return nil, err
		}
	}
	_db = &DB{
		DB:    db,
		dbCfg: dbCfg,
	}
	_db.getSqlxConn()
	return _db, nil
}

// CreateDB mysql 创建数据库
func CreateDB(DataSourceName string) (err error) {
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
		fmt.Println("createDB: sql.Open:", err)
		return err
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		fmt.Println("ping: ", source, " error: ", err)
		return err
	}
	_, err = db.Exec(
		fmt.Sprintf("CREATE DATABASE If Not Exists `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin",
			cfg.DBName))
	if err != nil {
		fmt.Println("createDB: db.exec:", err)
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
