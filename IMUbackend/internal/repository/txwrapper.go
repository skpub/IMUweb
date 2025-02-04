package repository

import (
	"IMUbackend/db"
	"context"
	"database/sql"
)

type TxManager interface {
	BeginTx(ctx context.Context) (Tx, error)
}

type Tx interface {
	Commit() error
	Rollback() error
	Queries() *db.Queries
}

type TxWrapper struct {
	tx      *sql.Tx
	queries *db.Queries
}

func (t *TxWrapper) Commit() error {
	return t.tx.Commit()
}

func (t *TxWrapper) Rollback() error {
	return t.tx.Rollback()
}

func (t *TxWrapper) Queries() *db.Queries {
	return t.queries
}

type SQLTxManager struct {
	db *sql.DB
}

func NewSQLTxManager(db *sql.DB) TxManager {
	return &SQLTxManager{db}
}

func (m *SQLTxManager) BeginTx(ctx context.Context) (Tx, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &TxWrapper{tx, db.New(tx)}, nil
}
