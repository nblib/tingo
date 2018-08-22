package tdbcomn

import (
	"context"
	"database/sql"
)

/**
封装sql的db的方法,去掉Close方法,防止关闭db
*/
func NewOperator(db *sql.DB) *DBOperator {
	return &DBOperator{db: db}
}

type DBOperator struct {
	db *sql.DB
}

func (d *DBOperator) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}
func (d *DBOperator) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}
func (d *DBOperator) Prepare(query string) (*sql.Stmt, error) {
	return d.db.Prepare(query)
}
func (d *DBOperator) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return d.db.PrepareContext(ctx, query)
}
func (d *DBOperator) Begin() (*sql.Tx, error) {
	return d.db.Begin()
}
func (d *DBOperator) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return d.db.BeginTx(ctx, opts)
}
func (d *DBOperator) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.Query(query, args...)
}

func (d *DBOperator) QueryRow(query string, args ...interface{}) *sql.Row {
	return d.db.QueryRow(query, args...)
}

func (d *DBOperator) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}
func (d *DBOperator) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}
func (d *DBOperator) Stats() sql.DBStats {
	return d.db.Stats()
}
