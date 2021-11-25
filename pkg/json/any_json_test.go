package anyjson

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyJsonDeserialisation(t *testing.T) {
	tests := []struct {
		description    string
		input          string
		expectedOutput AnyJson
	}{
		{
			description: "Deserialize float, int and string correctly",
			input: `{
				"string": "hello",
				"float": 3.5,
				"int": 3
			}`,
			expectedOutput: AnyJson{
				"string": "hello",
				"float":  3.5,
				"int":    int64(3),
			},
		},
		{
			description: "Deserialize negative float, int and string correctly",
			input: `{
				"string": "hello",
				"float": -3.5,
				"int": -3
			}`,
			expectedOutput: AnyJson{
				"string": "hello",
				"float":  -3.5,
				"int":    int64(-3),
			},
		},
		{
			description: "Deserialize float correctly to 3 decimal places",
			input: `{
				"float": 3.999
			}`,
			expectedOutput: AnyJson{
				"float": 3.999,
			},
		},
		{
			description: "Nested values",
			input: `{
				"float": 3.999,
				"meta": {
					"int": 76,
					"float": 5.2,
					"string": "world"
				}
			}`,
			expectedOutput: AnyJson{
				"float": 3.999,
				"meta": map[string]interface{}{
					"int":    int64(76),
					"float":  5.2,
					"string": "world",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			var a AnyJson

			json.Unmarshal([]byte(tt.input), &a)
			assert.Equal(t, tt.expectedOutput, a)
		})
	}
}
