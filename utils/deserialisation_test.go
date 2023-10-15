package utils

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	ErrorMessage string `json:"message"`
	ErrorCode    int    `json:"code"`
}

func (e testStruct) Error() string {
	return fmt.Sprintf("Error returned with reference %v and code %v", e.ErrorMessage, e.ErrorCode)
}

type errorTest2 struct {
	Info string `json:"info"`
}

func (e errorTest2) Error() string {
	return fmt.Sprintf("The test interface has info %v", e.Info)
}

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func TestDeserialiseRequest(t *testing.T) {

	var i testStruct
	o := &testStruct{
		ErrorMessage: "There is an error",
		ErrorCode:    400,
	}
	jsonObj := `{
		"message": "There is an error",
		"code": 400
	  }`
	rc := strings.NewReader(jsonObj)

	rcError := errReader(0)
	var iError errorTest2

	var tests = []struct {
		testName       string
		inputInterface interface{}
		inputBody      io.Reader
		expectedOutput interface{}
		expectedError  bool
	}{
		{testName: "successfully deserialise body into interface", inputInterface: &i, inputBody: rc, expectedOutput: o, expectedError: false},
		{testName: "returns error if ioutil.ReadAll fails", inputInterface: &i, inputBody: rcError, expectedOutput: nil, expectedError: true},
		{testName: "returns error if cannot unmarshal the json into the interface", inputInterface: &iError, inputBody: rc, expectedOutput: nil, expectedError: true},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {

			d, err := DeserialiseRequest(tt.inputInterface, tt.inputBody)

			if tt.expectedError == true {
				assert.Nil(t, d)
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, d)
			}
		})
	}
}
