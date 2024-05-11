package database_test

import (
	"Project1/database"
	"Project1/models"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntegrationAccounts_InsertAccounts(t *testing.T) {

	ctx := context.Background()

	testInfra := NewTestInfra()
	db := testInfra.GetDB()

	defer db.Close()

	accountCreated := models.Account{
		ID:      1,
		Balance: 100.0,
	}

	tests := []struct {
		name        string
		input       models.Accounts
		expectedErr bool
	}{
		{
			name:        "nil-accounts",
			input:       nil,
			expectedErr: true,
		},
		{
			name:        "success",
			input:       models.Accounts{accountCreated},
			expectedErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dbConn := database.NewDatabase(db)
			err := dbConn.InsertAccounts(ctx, tc.input)
			require.Equal(t, tc.expectedErr, err != nil)

			err = hardDeleteForLocalTesting(db)
			require.NoError(t, err)
		})
	}
}

func TestIntegrationAccounts_GetAccountsByID(t *testing.T) {

	ctx := context.Background()

	testInfra := NewTestInfra()
	db := testInfra.GetDB()

	defer db.Close()

	accountCreated := models.Account{
		ID:      1,
		Balance: 100.0,
	}

	tests := []struct {
		name         string
		input        int
		expectedResp models.Account
		expectedErr  bool
	}{
		{
			name:         "nothing to return",
			input:        0,
			expectedResp: models.Account{},
			expectedErr:  true,
		},
		{
			name:  "success",
			input: 1,
			expectedResp: models.Account{
				ID:      1,
				Balance: 100.0,
			},
			expectedErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dbConn := database.NewDatabase(db)

			err := dbConn.InsertAccounts(ctx, models.Accounts{accountCreated})
			require.NoError(t, err)

			gotten, err := dbConn.GetAccountsByID(ctx, tc.input)
			require.Equal(t, tc.expectedErr, err != nil)
			require.Equal(t, tc.expectedResp.ID, gotten.ID)
			require.Equal(t, tc.expectedResp.Balance, gotten.Balance)

			err = hardDeleteForLocalTesting(db)
			require.NoError(t, err)
		})
	}
}
