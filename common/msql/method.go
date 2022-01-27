package msql

import (
	"context"
	"database/sql"
)

func (_db *DB) Exec(format string, args ...interface{}) (sql.Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), _db.dbCfg.MaxQueryTime)
	defer cancel()
	return _db.DB.ExecContext(ctx, format, args...)
}

func (_db *DB) BeginTx(cb func(tx *sql.Tx) error, opts ...TxOption) error {
	defaultOpt := DefaultTxOptions()
	for _, opt := range opts {
		opt(defaultOpt)
	}
	ctx, cancel := context.WithTimeout(context.Background(), _db.dbCfg.MaxQueryTime)
	defer cancel()
	tx, err := _db.DB.BeginTx(ctx, defaultOpt)
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
	return nil
}
