package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: we currently deserialise the request twice for all calls
// 1. into a possible error interface
// 2. into the actual response body if there are no erros from req
// -> break into deserialise and unmarshal fns instead

// TODO: check the erros returned by the fn (added now)
// when it's called in CheckForErrorResponse and each endpt call
// error from ReadAll OR json.Unmarshal -> change to another error type we have made??

type errorTest struct {
	ErrorMessage string `json:"message"`
	ErrorCode    int    `json:"code"`
}

func (e errorTest) Error() string {
	return fmt.Sprintf("Error returned with reference %v and code %v", e.ErrorMessage, e.ErrorCode)
}

type bodyStruct struct {
	body bytes.Buffer
}

func (b bodyStruct) Close() error {
	return errorTest{
		ErrorMessage: "error",
		ErrorCode:    400,
	}
}
func (b bodyStruct) Read(p []byte) (int, error) {

	t, err := b.body.Read(p)

	if err != nil {
		return 0, err
	}

	return t, io.EOF
}

func TestDeserialiseRequest(t *testing.T) {

	var i1 errorTest

	o1 := &errorTest{
		ErrorMessage: "There is an error",
		ErrorCode:    400,
	}

	buf, _ := json.Marshal(o1)
	b := bytes.NewBuffer(buf)
	by := bodyStruct{
		body: *b,
	}

	var tests = []struct {
		testName       string
		inputInterface interface{}
		inputBody      io.ReadCloser
		expectedOutput interface{}
		expectedError  error
	}{
		{testName: "deserialise body into ?????", inputInterface: &i1, inputBody: by, expectedOutput: o1, expectedError: nil},
	}
	for _, tt := range tests {

		t.Run(tt.testName, func(t *testing.T) {

			d, err := DeserialiseRequest(tt.inputInterface, tt.inputBody)

			if err != nil {
				// add more error checking here
				assert.Equal(t, tt.expectedError, err)
			}

			assert.Equal(t, tt.expectedOutput, d)
		})
	}
}
