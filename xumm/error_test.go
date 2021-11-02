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
		assert.EqualValues(t, &ErrorResponse{ErrorResponseInternal: ErrorResponseInternal{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 812}}, err)
		assert.EqualError(t, err, "Error returned with reference 3a04c7d3-94aa-4d8d-9559-62bb5e8a653c and code 812")
	})
}
