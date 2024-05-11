package database_test

import (
	"Project1/database"
	"Project1/models"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntegrationTransactions_InsertTransactions(t *testing.T) {

	ctx := context.Background()

	testInfra := NewTestInfra()
	db := testInfra.GetDB()

	defer db.Close()

	accountCreated := models.Account{
		ID:      1,
		Balance: 100.0,
	}

	accountCreated2 := models.Account{
		ID:      2,
		Balance: 10.0,
	}

	tests := []struct {
		name        string
		input       models.Transactions
		expectedErr bool
	}{
		{
			name: "success",
			input: models.Transactions{
				{
					SourceAccountID:      1,
					DestinationAccountID: 2,
					Amount:               10.0,
				},
			},
			expectedErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dbConn := database.NewDatabase(db)
			tx, err := dbConn.Begin(ctx)
			require.NoError(t, err)

			defer tx.Rollback()

			err = dbConn.InsertAccounts(ctx, models.Accounts{accountCreated, accountCreated2})
			require.NoError(t, err)

			err = dbConn.InsertTransactions(ctx, tx, tc.input)
			require.Equal(t, tc.expectedErr, err != nil)

			err = tx.Commit()
			require.NoError(t, err)

			err = hardDeleteForLocalTesting(db)
			require.NoError(t, err)
		})
	}
}
