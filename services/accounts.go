package services

import (
	"Project1/database"
	"Project1/models"
	"context"
)

type AccountsInterface interface {
	GetAccount(ctx context.Context, ID int) (models.Account, error)
	PostAccount(ctx context.Context, req models.Account) error
}

type Accounts struct {
	database database.DatabaseInterface
}

func NewAccountsService(db database.DatabaseInterface) *Accounts {
	return &Accounts{
		database: db,
	}
}

func (s *Accounts) GetAccount(ctx context.Context, ID int) (models.Account, error) {
	return s.database.GetAccountsByID(ctx, ID)
}

func (s *Accounts) PostAccount(ctx context.Context, req models.Account) error {
	return s.database.InsertAccounts(ctx, []models.Account{req})
}
