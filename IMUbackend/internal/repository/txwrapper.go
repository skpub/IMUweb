package repository

import (
	"context"
	"IMUbackend/db"
)

type TxManager interface {
	BeginTx(ctx context.Context) (Tx, error)
}

type Tx interface {
	Commit() error
	Rollback() error
	Queries() *db.Queries
}
