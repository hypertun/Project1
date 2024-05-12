package models

import (
	"strconv"
	"time"
)

type TransactionHandlerReq struct {
	SourceAccountID      int    `db:"source_account_id" json:"source_account_id"`
	DestinationAccountID int    `db:"destination_account_id" json:"destination_account_id"`
	Amount               string `db:"amount" json:"amount"`
}

type Transaction struct {
	SourceAccountID      int       `db:"source_account_id" json:"source_account_id"`
	DestinationAccountID int       `db:"destination_account_id" json:"destination_account_id"`
	Amount               float64   `db:"amount" json:"amount"`
	Created_At           time.Time `db:"created_at"`
}

type Transactions []Transaction

func (t TransactionHandlerReq) ConvertToTransaction() (Transaction, error) {
	if t.SourceAccountID <= 0 || t.DestinationAccountID <= 0 {
		return Transaction{}, ErrInvalidAccountID
	}

	if t.SourceAccountID == t.DestinationAccountID {
		return Transaction{}, ErrSameAccount
	}

	amount, err := strconv.ParseFloat(t.Amount, 64)
	if err != nil {
		return Transaction{}, err
	}

	return Transaction{
		SourceAccountID:      t.SourceAccountID,
		DestinationAccountID: t.DestinationAccountID,
		Amount:               amount,
	}, nil
}
