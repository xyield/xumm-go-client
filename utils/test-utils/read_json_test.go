package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertFileToJsonString(t *testing.T) {
	tests := []struct {
		description    string
		inputFileName  string
		expectedOutput string
	}{
		{
			description:    "Test spaces in values aren't removed",
			inputFileName:  "static-test-data/jsonmap_with_spaces_in_values.json",
			expectedOutput: "{\"string\":\"Hello\",\"sentence\":\"Hello World, testing 1,2...\",\"int\":5}",
		},
		{
			description:    "Test nested objects in json map",
			inputFileName:  "static-test-data/jsonmap_with_nested_objects.json",
			expectedOutput: "{\"test\":\"testValue\",\"nested_object\":{\"sentence\":\"Testing 1,2,3...\",\"account\":\"testAccount\"}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			got := ConvertJsonFileToJsonString(tt.inputFileName)
			assert.JSONEq(t, tt.expectedOutput, got)
		})
	}
}
