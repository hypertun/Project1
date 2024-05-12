package models_test

import (
	"Project1/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModels_ConvertToTransaction(t *testing.T) {

	tests := []struct {
		name         string
		input        models.TransactionHandlerReq
		expectedResp models.Transaction
		expectedErr  bool
	}{
		{
			name: "invalidsourceid",
			input: models.TransactionHandlerReq{
				SourceAccountID:      -1,
				DestinationAccountID: 2,
				Amount:               "100",
			},
			expectedResp: models.Transaction{},
			expectedErr:  true,
		},
		{
			name: "invaliddestinationid",
			input: models.TransactionHandlerReq{
				SourceAccountID:      1,
				DestinationAccountID: -2,
				Amount:               "100",
			},
			expectedResp: models.Transaction{},
			expectedErr:  true,
		},
		{
			name: "sameaccount",
			input: models.TransactionHandlerReq{
				SourceAccountID:      1,
				DestinationAccountID: 1,
				Amount:               "100",
			},
			expectedResp: models.Transaction{},
			expectedErr:  true,
		},
		{
			name: "parsefloat",
			input: models.TransactionHandlerReq{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               "a",
			},
			expectedResp: models.Transaction{},
			expectedErr:  true,
		},
		{
			name: "success",
			input: models.TransactionHandlerReq{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               "1000",
			},
			expectedResp: models.Transaction{
				SourceAccountID:      1,
				DestinationAccountID: 2,
				Amount:               1000.0,
			},
			expectedErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := tc.input.ConvertToTransaction()
			require.Equal(t, tc.expectedErr, err != nil)
			assert.Equal(t, tc.expectedResp, resp)
		})
	}
}
