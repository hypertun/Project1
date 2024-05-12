package models

import "errors"

var (
	ErrInvalidAccountID = errors.New("account id cannot less than or equal to 0")
	ErrSameAccount      = errors.New("source and destination account cannot be same")

	//used in tests only
	ErrDummy = errors.New("dummy error")
)
