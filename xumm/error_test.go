// +build unit

package xumm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckForErrorResponse(t *testing.T) {

	tt := []struct {
		description    string
		json           string
		expectedOutput error
		httpStatus     int
	}{
		{
			description: "errorResponseReturnedWithoutMessage",
			json: `{
				"error": {
					"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
					"code": 812
				}
			}`,
			expectedOutput: &ErrorResponse{ErrorResponseBody: ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 812}},
			httpStatus:     403,
		},
		{
			description: "errorResponseReturnedWithMessage",
			json: `{
				"error": {
					"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
					"code": 812,
					"message": "Error message"
				}
			}`,
			expectedOutput: &ErrorResponse{ErrorResponseBody: ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Message: "Error message", Code: 812}},
			httpStatus:     403,
		},
		{
			description: "errorResponseReturnedWithMessageAndReference",
			json: `{
				"error": {
					"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
					"message": "Error message"
				}
			}`,

			expectedOutput: &ErrorResponse{ErrorResponseBody: ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Message: "Error message"}},
			httpStatus:     403,
		},
		{
			description: "errorUnauthorised",
			json: `{
				"error": true,
				"message": "message",
				"reference": "Endpoint unknown or method invalid for given endpoint",
				"code": 404,
				"req": "/v1/platform/payload/payload_uuid",
				"method": "GET"
			  }`,
			expectedOutput: &ErrorNotFound{
				Err:       true,
				Message:   "message",
				Reference: "Endpoint unknown or method invalid for given endpoint",
				Code:      404,
				Req:       "/v1/platform/payload/payload_uuid",
				Method:    "GET",
			},
			httpStatus: 404,
		},
		{
			description: "errorPartialUnauthorised",
			json: `{
				"reference": "Endpoint unknown or method invalid for given endpoint",
				"code": 404
			  }`,
			expectedOutput: &ErrorNotFound{
				Reference: "Endpoint unknown or method invalid for given endpoint",
				Code:      404,
			},
			httpStatus: 404,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			b := ioutil.NopCloser(bytes.NewReader([]byte(test.json)))
			res := &http.Response{
				StatusCode: test.httpStatus,
				Body:       b,
			}
			err := CheckForErrorResponse(res)
			assert.Error(t, err)
			assert.EqualValues(t, test.expectedOutput, err)
			assert.EqualError(t, err, test.expectedOutput.Error())
		})
	}

}
