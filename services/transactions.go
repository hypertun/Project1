package services

import (
	"Project1/database"
	"Project1/models"
	"context"
)

type TransactionsInterface interface {
	PostTransaction(ctx context.Context, req models.Transaction) error
}

type Transactions struct {
	database database.DatabaseInterface
}

func NewTransactionsService(db database.DatabaseInterface) *Transactions {
	return &Transactions{
		database: db,
	}
}

func (t *Transactions) PostTransaction(ctx context.Context, req models.Transaction) error {
	tx, err := t.database.Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := t.database.InsertTransactions(ctx, tx, []models.Transaction{req}); err != nil {
		return err
	}

	sourceAccount := models.Account{
		ID:      req.SourceAccountID,
		Balance: -req.Amount,
	}

	destinationAccount := models.Account{
		ID:      req.DestinationAccountID,
		Balance: req.Amount,
	}

	if err := t.database.UpdateAccounts(ctx, tx, sourceAccount); err != nil {
		return err
	}

	if err := t.database.UpdateAccounts(ctx, tx, destinationAccount); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
