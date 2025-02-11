package infrastructure

import (
	"IMUbackend/db"
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type IDBTX interface {
	BeginTx(ctx context.Context, ops *sql.TxOptions) (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
	GetDBTX() db.DBTX
}

type DB struct {
	db *sql.DB
}

func (d *DB) BeginTx(ctx context.Context, ops *sql.TxOptions) (*sql.Tx, error) {
	tx, err := d.db.BeginTx(ctx, ops)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (d *DB) GetDBTX() db.DBTX {
	return d.db
}

func (d *DB) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (d *DB) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func NewDBConnection(dsn string) (IDBTX, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
