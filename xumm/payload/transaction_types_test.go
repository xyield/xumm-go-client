package payload

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionTypeToString(t *testing.T) {
	tests := []struct {
		description     string
		transactionType TransactionType
		expectedString  string
	}{
		{
			description:     "Check SignIn transaction type returns correct string representation",
			transactionType: SignIn,
			expectedString:  "SignIn",
		},
		{
			description:     "Check payment has correct string",
			transactionType: Payment,
			expectedString:  "Payment",
		},
		{
			description:     "Check enable amendment has correct string",
			transactionType: EnableAmendment,
			expectedString:  "EnableAmendment",
		},
		{
			description:     "Check transaction type outside range",
			transactionType: TransactionType(100),
			expectedString:  "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			assert.Equal(t, tt.expectedString, tt.transactionType.String())
		})
	}
}

func TestTransactionTypeDeserialisation(t *testing.T) {
	tests := []struct {
		description    string
		input          []byte
		expectedOutput transactionType
		expectedError  error
	}{
		{
			description: "Valid transaction type",
			input:       json.RawMessage(`{"TransactionType":"Payment"}`),
			expectedOutput: transactionType{
				Type: Payment,
			},
			expectedError: nil,
		},
		{
			description: "Invalid transaction type",
			input:       json.RawMessage(`{"TransactionType":"NotValid"}`),
			expectedOutput: transactionType{
				Type: TransactionType(0),
			},
			expectedError: &TransactionTypeError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			var actual transactionType
			err := json.Unmarshal(tt.input, &actual)

			if tt.expectedError != nil {
				assert.Error(t, tt.expectedError)
				assert.EqualError(t, err, tt.expectedError.Error())
				assert.Equal(t, tt.expectedOutput, actual)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, actual)
			}
		})
	}
}
