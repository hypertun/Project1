package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	ErrInvalidAccountID = errors.New("account id cannot less than or equal to 0")
)

type AccountCreationRequest struct {
	AccountID      int    `json:"account_id"`
	InitialBalance string `json:"initial_balance"`
}

func (a AccountCreationRequest) ConvertToAccount() (Account, error) {
	if a.AccountID <= 0 {
		return Account{}, ErrInvalidAccountID
	}

	balance, err := strconv.ParseFloat(a.InitialBalance, 64)
	if err != nil {
		return Account{}, err
	}

	return Account{
		ID:      a.AccountID,
		Balance: balance,
	}, nil
}

type Account struct {
	ID         int       `db:"id" json:"account_id"`
	Balance    float64   `db:"balance" json:"balance"`
	Created_At time.Time `db:"created_at"`
	Updated_At time.Time `db:"updated_at"`
}

type Accounts []Account
