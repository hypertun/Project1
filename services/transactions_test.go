package services_test

import (
	"Project1/database/mocks"
	"Project1/models"
	"Project1/services"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTransactionsService_PostTransaction(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                      string
		expectedDatabaseInterface func() (*mocks.DatabaseInterface, sqlmock.Sqlmock)
		input                     models.Transaction
		expectedErr               bool
	}{
		{
			name: "begin",
			expectedDatabaseInterface: func() (*mocks.DatabaseInterface, sqlmock.Sqlmock) {
				mockDB, mockExt, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				require.NoError(t, err)

				defer mockDB.Close()

				db := mocks.NewDatabaseInterface(t)

				db.On("Begin", ctx).Return(nil, models.ErrDummy)

				return db, mockExt
			},
			input: models.Transaction{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               100,
			},
			expectedErr: true,
		},
		{
			name: "inserterror",
			expectedDatabaseInterface: func() (*mocks.DatabaseInterface, sqlmock.Sqlmock) {
				mockDB, mockExt, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				require.NoError(t, err)

				defer mockDB.Close()

				sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

				mockExt.ExpectBegin()

				sqlxTx, err := sqlxDB.BeginTxx(ctx, nil)
				require.NoError(t, err)

				db := mocks.NewDatabaseInterface(t)

				db.On("Begin", ctx).Return(sqlxTx, nil)

				db.On("InsertTransactions", ctx, sqlxTx, models.Transactions{
					{
						SourceAccountID:      1,
						DestinationAccountID: 2,
						Amount:               100,
					},
				}).Return(models.ErrDummy)

				mockExt.ExpectRollback()

				return db, mockExt
			},
			input: models.Transaction{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               100,
			},
			expectedErr: true,
		},
		{
			name: "updateSourceAccounterror",
			expectedDatabaseInterface: func() (*mocks.DatabaseInterface, sqlmock.Sqlmock) {
				mockDB, mockExt, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				require.NoError(t, err)

				defer mockDB.Close()

				sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

				mockExt.ExpectBegin()

				sqlxTx, err := sqlxDB.BeginTxx(ctx, nil)
				require.NoError(t, err)

				db := mocks.NewDatabaseInterface(t)

				db.On("Begin", ctx).Return(sqlxTx, nil)

				db.On("InsertTransactions", ctx, sqlxTx, models.Transactions{
					{
						SourceAccountID:      1,
						DestinationAccountID: 2,
						Amount:               100,
					},
				}).Return(nil)

				db.On("UpdateAccounts", ctx, sqlxTx, models.Account{
					ID:      1,
					Balance: -100,
				}).Return(models.ErrDummy)

				mockExt.ExpectRollback()

				return db, mockExt
			},
			input: models.Transaction{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               100,
			},
			expectedErr: true,
		},
		{
			name: "updateDestinationAccountError",
			expectedDatabaseInterface: func() (*mocks.DatabaseInterface, sqlmock.Sqlmock) {
				mockDB, mockExt, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				require.NoError(t, err)

				defer mockDB.Close()

				sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

				mockExt.ExpectBegin()

				sqlxTx, err := sqlxDB.BeginTxx(ctx, nil)
				require.NoError(t, err)

				db := mocks.NewDatabaseInterface(t)

				db.On("Begin", ctx).Return(sqlxTx, nil)

				db.On("InsertTransactions", ctx, sqlxTx, models.Transactions{
					{
						SourceAccountID:      1,
						DestinationAccountID: 2,
						Amount:               100,
					},
				}).Return(nil)

				db.On("UpdateAccounts", ctx, sqlxTx, models.Account{
					ID:      1,
					Balance: -100,
				}).Return(nil)

				db.On("UpdateAccounts", ctx, sqlxTx, models.Account{
					ID:      2,
					Balance: 100,
				}).Return(models.ErrDummy)

				mockExt.ExpectRollback()

				return db, mockExt
			},
			input: models.Transaction{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               100,
			},
			expectedErr: true,
		},
		{
			name: "commiterr",
			expectedDatabaseInterface: func() (*mocks.DatabaseInterface, sqlmock.Sqlmock) {
				mockDB, mockExt, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				require.NoError(t, err)

				defer mockDB.Close()

				sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

				mockExt.ExpectBegin()

				sqlxTx, err := sqlxDB.BeginTxx(ctx, nil)
				require.NoError(t, err)

				db := mocks.NewDatabaseInterface(t)

				db.On("Begin", ctx).Return(sqlxTx, nil)

				db.On("InsertTransactions", ctx, sqlxTx, models.Transactions{
					{
						SourceAccountID:      1,
						DestinationAccountID: 2,
						Amount:               100,
					},
				}).Return(nil)

				db.On("UpdateAccounts", ctx, sqlxTx, models.Account{
					ID:      1,
					Balance: -100,
				}).Return(nil)

				db.On("UpdateAccounts", ctx, sqlxTx, models.Account{
					ID:      2,
					Balance: 100,
				}).Return(nil)

				mockExt.ExpectCommit().WillReturnError(models.ErrDummy)

				return db, mockExt
			},
			input: models.Transaction{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               100,
			},
			expectedErr: true,
		},
		{
			name: "success",
			expectedDatabaseInterface: func() (*mocks.DatabaseInterface, sqlmock.Sqlmock) {
				mockDB, mockExt, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
				require.NoError(t, err)

				defer mockDB.Close()

				sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

				mockExt.ExpectBegin()

				sqlxTx, err := sqlxDB.BeginTxx(ctx, nil)
				require.NoError(t, err)

				db := mocks.NewDatabaseInterface(t)

				db.On("Begin", ctx).Return(sqlxTx, nil)

				db.On("InsertTransactions", ctx, sqlxTx, models.Transactions{
					{
						SourceAccountID:      1,
						DestinationAccountID: 2,
						Amount:               100,
					},
				}).Return(nil)

				db.On("UpdateAccounts", ctx, sqlxTx, models.Account{
					ID:      1,
					Balance: -100,
				}).Return(nil)

				db.On("UpdateAccounts", ctx, sqlxTx, models.Account{
					ID:      2,
					Balance: 100,
				}).Return(nil)

				mockExt.ExpectCommit()

				return db, mockExt
			},
			input: models.Transaction{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               100,
			},
			expectedErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB, mockExt := tc.expectedDatabaseInterface()

			err := services.NewTransactionsService(mockDB).PostTransaction(ctx, tc.input)

			require.Equal(t, tc.expectedErr, err != nil)

			err = mockExt.ExpectationsWereMet()
			require.NoError(t, err)

			mock.AssertExpectationsForObjects(t, mockDB)
		})
	}
}
