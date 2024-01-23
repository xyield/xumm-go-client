//go:build unit
// +build unit

package meta

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetHookhash(t *testing.T) {

	c := &models.HookHashResponse{
		Name:        "testhookhash",
		Description: "atesthookhash",
		Creator: models.HookHashCreator{
			Name: "joebloggs",
			Mail: "jb@gfail.com",
			Site: "test.com",
		},
		Xapp:             "testXapp",
		AppUUID:          "testUUID",
		Icon:             "testIcon",
		VerifiedAccounts: []string{"testva"},
		Audits:           []string{"testaudits"},
	}

	validJson := `{
		"name": "testhookhash",
		"description": "atesthookhash",
		"creator": {
			"name": "joebloggs",
			"mail": "jb@gfail.com",
			"site": "test.com"
		},
		"xapp": "testXapp",
		"appuuid": "testUUID",
		"icon": "testIcon",
		"verifiedaccounts": ["testva"],
		"audits": ["testaudits"]
		}`

	errorJson := `{
		"error": {
			"reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
			"code": 400
			"message": "Invalid hookhash provided, must be 64 hexadecimal characters"
		}
	}`

	var tests = []struct {
		testName       string
		testValue      string
		inputValue     string
		expectedOutput *models.HookHashResponse
		expectedError  error
		httpStatusCode int
	}{
		{
			testName:       "GetHookhash - success",
			testValue:      "A17E4DEAD62BF705D9B73B4EAD2832F1C55C6C5A0067327A45E497FD8D31C0E3",
			inputValue:     validJson,
			expectedOutput: c,
			expectedError:  nil,
			httpStatusCode: http.StatusOK,
		},
		{
			testName:       "GetHookhash - error response",
			testValue:      "invalidhookhash",
			inputValue:     errorJson,
			expectedOutput: nil,
			expectedError:  &InvalidHookHash{},
			httpStatusCode: 400,
		},
	}
	for _, tt := range tests {

		t.Run(tt.testName, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.inputValue, tt.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			meta := &Meta{
				Cfg: cfg,
			}
			ca, err := meta.GetHookHash(tt.testValue)

			if tt.expectedError != nil {
				assert.Nil(t, ca)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, ca)
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
			}
		})
	}
}
