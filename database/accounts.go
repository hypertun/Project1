package database

import (
	"Project1/models"
	"context"

	"github.com/jmoiron/sqlx"
)

const (
	insertAccountsStatement = `INSERT INTO accounts (id, balance) VALUES (:id, :balance) ON CONFLICT (id) DO NOTHING`
	getAccountStatement     = `SELECT id, balance, created_at, updated_at FROM accounts where id = $1`
	updateAccountsStatement = `UPDATE accounts SET balance = balance + :balance WHERE id = :id`
)

func (d *Database) InsertAccounts(ctx context.Context, req models.Accounts) error {
	if _, err := d.db.NamedExecContext(ctx, insertAccountsStatement, req); err != nil {
		return err
	}

	return nil
}

func (d *Database) GetAccountsByID(ctx context.Context, ID int) (models.Account, error) {
	var account models.Account

	if err := d.db.GetContext(ctx, &account, getAccountStatement, ID); err != nil {
		return account, err
	}

	return account, nil
}

func (d *Database) UpdateAccounts(ctx context.Context, tx *sqlx.Tx, req models.Account) error {
	if _, err := tx.NamedExecContext(ctx, updateAccountsStatement, req); err != nil {
		return err
	}

	return nil
}
