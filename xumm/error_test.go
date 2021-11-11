package xumm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckForError(t *testing.T) {
	t.Run("If error returns function creates and logs an error", func(t *testing.T) {
		json := `{
			"error": {
				"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
				"code": 812
			}
		}`
		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: http.StatusForbidden,
			Body:       b,
		}
		err := CheckForErrorResponse(res)
		assert.Error(t, err)
		assert.EqualValues(t, &ErrorResponse{ErrorResponseBody: ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 812}}, err)
		assert.EqualError(t, err, "Error returned with reference 3a04c7d3-94aa-4d8d-9559-62bb5e8a653c and code 812")
	})

	t.Run("If error response returns with messsage, function creates and logs an error", func(t *testing.T) {
		json := `{
			"error": {
				"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
				"code": 812,
				"message": "Error message"
			}
		}`
		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: http.StatusForbidden,
			Body:       b,
		}
		err := CheckForErrorResponse(res)
		assert.Error(t, err)
		assert.EqualValues(t, &ErrorResponse{ErrorResponseBody: ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Message: "Error message", Code: 812}}, err)
		assert.EqualError(t, err, "Error returned with reference 3a04c7d3-94aa-4d8d-9559-62bb5e8a653c, code 812 and message 'Error message'")
	})

	t.Run("If error response returns with messsage and reference, function creates and logs an error", func(t *testing.T) {
		json := `{
			"error": {
				"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
				"message": "Error message"
			}
		}`
		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: http.StatusForbidden,
			Body:       b,
		}
		err := CheckForErrorResponse(res)
		assert.Error(t, err)
		assert.EqualValues(t, &ErrorResponse{ErrorResponseBody: ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Message: "Error message"}}, err)
		assert.EqualError(t, err, "Error returned with reference 3a04c7d3-94aa-4d8d-9559-62bb5e8a653c and message 'Error message'")
	})

	t.Run("If unauthorised error returns function creates and logs an error", func(t *testing.T) {

		expectedError := &ErrorNotFound{
			Reference: "Endpoint unknown or method invalid for given endpoint",
			Code:      404,
			Message:   "message",
			Req:       "/v1/platform/payload/payload_uuid",
			Method:    "GET",
			Err:       true,
		}

		json := `{
			"error": true,
			"message": "message",
			"reference": "Endpoint unknown or method invalid for given endpoint",
			"code": 404,
			"req": "/v1/platform/payload/payload_uuid",
			"method": "GET"
		  }`

		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       b,
		}
		err := CheckForErrorResponse(res)
		assert.Error(t, err)
		assert.EqualValues(t, expectedError, err)
		assert.EqualError(t, err, "Error returned with code 404, reference 'Endpoint unknown or method invalid for given endpoint' and message 'message'")
	})
	t.Run("If partial unauthorised error returns function creates and logs an error", func(t *testing.T) {

		expectedError := &ErrorNotFound{
			Reference: "Endpoint unknown or method invalid for given endpoint",
			Code:      404,
		}
		json := `{
			"reference": "Endpoint unknown or method invalid for given endpoint",
			"code": 404
		  }`

		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: http.StatusNotFound,
			Body:       b,
		}
		err := CheckForErrorResponse(res)
		assert.Error(t, err)
		assert.EqualValues(t, expectedError, err)
		assert.EqualError(t, err, "Error returned with code 404 and reference 'Endpoint unknown or method invalid for given endpoint'")
	})
}
