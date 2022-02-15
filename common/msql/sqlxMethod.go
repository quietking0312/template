package msql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func (_db *DB) GetSqlxConn() *sqlx.DB {
	return _db.SqlxDB
}

func (_db *DB) getSqlxConn() {
	_db.SqlxDB = sqlx.NewDb(_db.DB, _db.dbCfg.DriveName)
	return
}

func GetSqlxConn() *sqlx.DB {
	return _db.GetSqlxConn()
}

func (_db *DB) SqlxBeginTx(cb func(tx *sqlx.Tx, ctx context.Context) error, opts ...TxOption) error {
	defaultOpt := DefaultTxOptions()
	for _, opt := range opts {
		opt(defaultOpt)
	}
	ctx, cancel := context.WithTimeout(context.Background(), _db.dbCfg.MaxQueryTime)
	defer cancel()
	tx, err := _db.SqlxDB.BeginTxx(ctx, defaultOpt)
	if err != nil {
		return err
	}
	if err := cb(tx, ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return err
}

func (_db *DB) SqlxNameExec(format string, arg interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), _db.dbCfg.MaxQueryTime)
	defer cancel()
	return _db.SqlxDB.NamedExecContext(ctx, format, arg)
}

func (_db *DB) SqlxExec(format string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), _db.dbCfg.MaxQueryTime)
	defer cancel()
	return _db.SqlxDB.ExecContext(ctx, format, args...)
}

func (_db *DB) SqlxNameQuery(format string, args interface{}, cb func(rows *sqlx.Rows) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), _db.dbCfg.MaxQueryTime)
	defer cancel()
	rows, err := _db.SqlxDB.NamedQueryContext(ctx, format, args)
	if err != nil {
		return err
	}
	defer rows.Close()
	err = cb(rows)
	return err
}

func (_db *DB) SqlxGet(dest interface{}, format string, args ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), _db.dbCfg.MaxQueryTime)
	defer cancel()
	return _db.SqlxDB.GetContext(ctx, dest, format, args...)
}

func (_db *DB) SqlxSelect(dest interface{}, format string, args ...interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), _db.dbCfg.MaxQueryTime)
	defer cancel()
	return _db.SqlxDB.SelectContext(ctx, dest, format, args...)
}

func (_db *DB) In(format string, args ...interface{}) (string, []interface{}, error) {
	return sqlx.In(format, args...)
}
