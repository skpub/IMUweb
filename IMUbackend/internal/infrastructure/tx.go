package infrastructure

import (
	repo "IMUbackend/internal/repository"
	"context"
	"database/sql"
	"IMUbackend/db"
)


type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) repo.TxManager {
	return &DBManager{db}
}

func (m *DBManager) BeginTx(ctx context.Context) (repo.Tx, error) {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &TxWrapper{tx, db.New(tx)}, nil
}

type TxWrapper struct {
	tx 		*sql.Tx
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
