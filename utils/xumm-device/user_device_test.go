package xummdevice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateBearerToken(t *testing.T) {
	ud := &UserDevice{
		AccessToken:            "testToken",
		UniqueDeviceIdentifier: "testUniqueIdentifier",
	}
	tt := []struct {
		description string
		userDevice  *UserDevice
		input       string
		expected    string
	}{
		{
			description: "Test basic functionality",
			userDevice:  ud,
			input:       "1644938305000",
			expected:    "testToken.1644938305000.f6e855f62991033b5c035bc128f47493ff90c362182aafed3dfcf91df6a8093a",
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			got := tc.userDevice.generateBearerToken(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
