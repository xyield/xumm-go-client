// +build unit

package payload

import (
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
