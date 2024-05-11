package database

import (
	"Project1/models"
	"context"

	"github.com/jmoiron/sqlx"
)

const (
	insertTransactionsStatement = `INSERT INTO transactions (source_account_id, destination_account_id, amount) VALUES (:source_account_id, :destination_account_id, :amount)`
)

func (d *Database) InsertTransactions(ctx context.Context, tx *sqlx.Tx, req models.Transactions) error {
	if _, err := tx.NamedExecContext(ctx, insertTransactionsStatement, req); err != nil {
		return err
	}

	return nil
}
