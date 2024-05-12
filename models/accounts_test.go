package models_test

import (
	"Project1/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModels_ConvertToAccount(t *testing.T) {

	tests := []struct {
		name         string
		input        models.AccountCreationRequest
		expectedResp models.Account
		expectedErr  bool
	}{
		{
			name: "invalidaccountid",
			input: models.AccountCreationRequest{
				AccountID:      -1,
				InitialBalance: "100",
			},
			expectedResp: models.Account{},
			expectedErr:  true,
		},
		{
			name: "parsefloat",
			input: models.AccountCreationRequest{
				AccountID:      1,
				InitialBalance: "a",
			},
			expectedResp: models.Account{},
			expectedErr:  true,
		},
		{
			name: "success",
			input: models.AccountCreationRequest{
				AccountID:      1,
				InitialBalance: "100",
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
			resp, err := tc.input.ConvertToAccount()
			require.Equal(t, tc.expectedErr, err != nil)
			assert.Equal(t, tc.expectedResp, resp)
		})
	}
}
