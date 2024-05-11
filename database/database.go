package database

import (
	"Project1/models"
	"context"

	"github.com/jmoiron/sqlx"
)

type DatabaseInterface interface {
	Begin(ctx context.Context) (*sqlx.Tx, error)
	InsertAccounts(ctx context.Context, req models.Accounts) error
	GetAccountsByID(ctx context.Context, ID int) (models.Account, error)
	UpdateAccounts(ctx context.Context, tx *sqlx.Tx, req models.Account) error
	InsertTransactions(ctx context.Context, tx *sqlx.Tx, req models.Transactions) error
}

type Database struct {
	db *sqlx.DB
}

func NewDatabase(db *sqlx.DB) DatabaseInterface {
	return &Database{
		db: db,
	}
}

func (d *Database) Begin(ctx context.Context) (*sqlx.Tx, error) {
	return d.db.BeginTxx(ctx, nil)
}
