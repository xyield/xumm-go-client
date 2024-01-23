package meta

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	statictestdata "github.com/xyield/xumm-go-client/xumm/meta/static-test-data"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetAccountMeta(t *testing.T) {

	jsonResponse := `{
		"account": "rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ",
		"kycApproved": true,
		"xummPro": true,
		"blocked": false,
		"force_dtag": false,
		"avatar": "https://xumm.app/avatar/rwietsevLFg8XSmG3bEZzFein1g8RBqWDZ.png",
		"xummProfile": {
		  "accountAlias": "XRPL Labs - Wietse Wind",
		  "ownerAlias": "Wietse Wind",
		  "slug": "wietsewind",
		  "profileUrl": "https://xumm.me/wietsewind",
		  "accountSlug": null,
		  "payString": "wietsewind$xumm.me"
		},
		"thirdPartyProfiles": [
		  {
			"accountAlias": "Wietse Wind",
			"source": "xumm.app"
		  },
		  {
			"accountAlias": "wietse.com",
			"source": "xrpl"
		  },
		  {
			"accountAlias": "XRPL-Labs",
			"source": "bithomp.com"
		  }
		],
		"globalid": {
		  "linked": "2021-06-29T10:22:25.000Z",
		  "profileUrl": "https://app.global.id/u/wietse",
		  "sufficientTrust": true
		}
	  }`

	tests := []struct {
		testName       string
		inputValue     string
		expectedOutput *models.AccountMetaResponse
		expectedError  error
		httpStatusCode int
	}{
		{testName: "correctData", inputValue: jsonResponse, expectedOutput: statictestdata.AccountMetaTestResult, expectedError: nil, httpStatusCode: 200},
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
			ca, err := meta.GetAccountMeta(statictestdata.AccountMetaTestResult.Account)

			if tt.expectedError != nil {
				assert.Nil(t, ca)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, ca)
			}
		})

	}

}
