package msql

import (
	"context"
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

func (_db *DB) SqlxBeginTx(cb func(tx *sqlx.Tx) error, opts ...TxOption) error {
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
	if err := cb(tx); err != nil {
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
