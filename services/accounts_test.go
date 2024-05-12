package services_test

import (
	"Project1/database/mocks"
	"Project1/models"
	"Project1/services"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAccountsService_GetAccount(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                      string
		expectedDatabaseInterface func() *mocks.DatabaseInterface
		expectedResp              models.Account
		expectedErr               bool
	}{
		{
			name: "error",
			expectedDatabaseInterface: func() *mocks.DatabaseInterface {
				db := mocks.NewDatabaseInterface(t)
				db.On("GetAccountsByID", ctx, 1).Return(models.Account{
					ID:      1,
					Balance: 100.0,
				}, models.ErrDummy)
				return db
			},
			expectedResp: models.Account{
				ID:      1,
				Balance: 100.0,
			},
			expectedErr: true,
		},
		{
			name: "success",
			expectedDatabaseInterface: func() *mocks.DatabaseInterface {
				db := mocks.NewDatabaseInterface(t)
				db.On("GetAccountsByID", ctx, 1).Return(models.Account{
					ID:      1,
					Balance: 100.0,
				}, nil)
				return db
			},
			expectedResp: models.Account{
				ID:      1,
				Balance: 100.0,
			},
			expectedErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := tc.expectedDatabaseInterface()
			gotten, err := services.NewAccountsService(mockDB).GetAccount(ctx, 1)

			require.Equal(t, tc.expectedResp, gotten)
			require.Equal(t, tc.expectedErr, err != nil)

			mock.AssertExpectationsForObjects(t, mockDB)
		})
	}
}

func TestAccountsService_PostAccount(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name                      string
		expectedDatabaseInterface func() *mocks.DatabaseInterface
		input                     models.Account
		expectedErr               bool
	}{
		{
			name: "error",
			expectedDatabaseInterface: func() *mocks.DatabaseInterface {
				db := mocks.NewDatabaseInterface(t)
				db.On("InsertAccounts", ctx, models.Accounts{
					{
						ID:      1,
						Balance: 100.0,
					},
				}).Return(models.ErrDummy)
				return db
			},
			input: models.Account{
				ID:      1,
				Balance: 100.0,
			},
			expectedErr: true,
		},
		{
			name: "success",
			expectedDatabaseInterface: func() *mocks.DatabaseInterface {
				db := mocks.NewDatabaseInterface(t)
				db.On("InsertAccounts", ctx, models.Accounts{
					{
						ID:      1,
						Balance: 100.0,
					},
				}).Return(nil)
				return db
			},
			input: models.Account{
				ID:      1,
				Balance: 100.0,
			},
			expectedErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := tc.expectedDatabaseInterface()
			err := services.NewAccountsService(mockDB).PostAccount(ctx, tc.input)

			require.Equal(t, tc.expectedErr, err != nil)

			mock.AssertExpectationsForObjects(t, mockDB)
		})
	}
}
