package meta

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestVerifyUserToken(t *testing.T) {
	tt := []struct {
		description    string
		input          string
		json           string
		expectedOutput *models.UserTokenResponse
		expectedError  error
		httpStatusCode int
	}{
		{
			description: "Happy path",
			input:       "7e5d2547-4257-4487-afab-4a94bf07e92e",
			expectedOutput: &models.UserTokenResponse{
				Tokens: []models.UserTokenValidity{
					{
						UserToken: "7e5d2547-4257-4487-afab-4a94bf07e92e",
						Active:    true,
						Issued:    int64(1646767028),
						Expires:   int64(1650572328),
					},
				},
			},
			json: `{
				"tokens" : [
					{
						"user_token": "7e5d2547-4257-4487-afab-4a94bf07e92e",
						"active" : true,
						"issued": 1646767028,
						"expires": 1650572328
					}
				]
			}`,
			expectedError:  nil,
			httpStatusCode: 200,
		},
		{
			description:    "Empty user token string returns error",
			input:          "",
			json:           ``,
			expectedOutput: nil,
			expectedError:  &EmptyUserToken{},
			httpStatusCode: 400,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tc.json, tc.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			meta := &Meta{
				Cfg: cfg,
			}

			ut, err := meta.VerifyUserToken(tc.input)
			if tc.expectedError != nil {
				assert.Nil(t, ut)
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, ut)
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
			}
		})
	}
}

func TestVerifyUserTokens(t *testing.T) {
	tt := []struct {
		description     string
		input           []string
		json            string
		expectedOutput  *models.UserTokenResponse
		expectedReqBody string
		expectedError   error
		httpStatusCode  int
	}{
		{
			description: "Single token",
			input:       []string{"7e5d2547-4257-4487-afab-4a94bf07e92e"},
			expectedReqBody: `{
				"tokens" : [
					"7e5d2547-4257-4487-afab-4a94bf07e92e"
				]
			}`,
			expectedOutput: &models.UserTokenResponse{
				Tokens: []models.UserTokenValidity{
					{
						UserToken: "7e5d2547-4257-4487-afab-4a94bf07e92e",
						Active:    true,
						Issued:    int64(1646767028),
						Expires:   int64(1650572328),
					},
				},
			},
			json: `{
				"tokens" : [
					{
						"user_token": "7e5d2547-4257-4487-afab-4a94bf07e92e",
						"active" : true,
						"issued": 1646767028,
						"expires": 1650572328
					}
				]
			}`,
			expectedError:  nil,
			httpStatusCode: 200,
		},
		{
			description: "Multiple tokens",
			input: []string{
				"7e5d2547-4257-4487-afab-4a94bf07e92e",
				"0f14e888-7a35-449e-8fc7-1edd04c687ef",
				"d72f00e9-bb11-4fa4-9c06-e19611069f17",
			},
			expectedReqBody: `{
				"tokens" : [
					"7e5d2547-4257-4487-afab-4a94bf07e92e",
					"0f14e888-7a35-449e-8fc7-1edd04c687ef",
					"d72f00e9-bb11-4fa4-9c06-e19611069f17"
				]
			}`,
			expectedOutput: &models.UserTokenResponse{
				Tokens: []models.UserTokenValidity{
					{
						UserToken: "7e5d2547-4257-4487-afab-4a94bf07e92e",
						Active:    true,
						Issued:    int64(1646767028),
						Expires:   int64(1650572328),
					},
					{
						UserToken: "0f14e888-7a35-449e-8fc7-1edd04c687ef",
						Active:    true,
						Issued:    int64(1646767028),
						Expires:   int64(1650572328),
					},
					{
						UserToken: "d72f00e9-bb11-4fa4-9c06-e19611069f17",
						Active:    true,
						Issued:    int64(1646767028),
						Expires:   int64(1650572328),
					},
				},
			},
			json: `{
				"tokens" : [
					{
						"user_token": "7e5d2547-4257-4487-afab-4a94bf07e92e",
						"active" : true,
						"issued": 1646767028,
						"expires": 1650572328
					},
					{
						"user_token": "0f14e888-7a35-449e-8fc7-1edd04c687ef",
						"active" : true,
						"issued": 1646767028,
						"expires": 1650572328
					},
					{
						"user_token": "d72f00e9-bb11-4fa4-9c06-e19611069f17",
						"active" : true,
						"issued": 1646767028,
						"expires": 1650572328
					}
				]
			}`,
			expectedError:  nil,
			httpStatusCode: 200,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {

			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tc.json, tc.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			meta := &Meta{
				Cfg: cfg,
			}

			ut, err := meta.VerifyUserTokens(tc.input...)
			if tc.expectedError != nil {
				assert.Nil(t, ut)
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, ut)
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
				reqBody, _ := ioutil.ReadAll(m.Spy.Body)
				assert.JSONEq(t, tc.expectedReqBody, string(reqBody))
			}
		})
	}
}
