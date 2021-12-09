package xapp

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/pkg/json"
	testutils "github.com/xyield/xumm-go-client/pkg/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestPostXappPush(t *testing.T) {

	var tests = []struct {
		description    string
		request        models.XappRequest
		jsonRequest    string
		jsonResponse   string
		expectedOutput *models.XappResponse
		expectedError  error
		httpStatusCode int
	}{
		{
			description: "successfully create push",
			request: models.XappRequest{
				UserToken: "token",
				Subtitle:  "subtitle",
				Body:      "body",
				Data:      anyjson.AnyJson{},
			},
			jsonRequest: `{
				"user_token": "token",
				"subtitle": "subtitle",
				"body": "body"
			}`,
			jsonResponse: `{
				"pushed": true
			  }`,
			expectedOutput: &models.XappResponse{Pushed: true},
			expectedError:  nil,
			httpStatusCode: 200,
		},
		{
			description: "invalid request body",
			request: models.XappRequest{
				UserToken: "",
				Subtitle:  "",
				Body:      "body",
				Data:      anyjson.AnyJson{"test_json": "TestJson"},
			},
			jsonRequest:    "",
			jsonResponse:   "",
			expectedOutput: nil,
			expectedError:  &invalidPushRequestError{},
			httpStatusCode: 200,
		},
		{
			description: "error creating push",
			request: models.XappRequest{
				UserToken: "token",
				Subtitle:  "subtitle",
				Body:      "body",
				Data:      anyjson.AnyJson{"test_json": "TestJson"},
			},
			jsonRequest: `{
				"user_token": "token",
				"subtitle": "subtitle",
				"body": "body"
			}`,
			jsonResponse: `{
				"error": {
				  "reference": "42d58b17-ee92-419d-b8ec-15797d10c4ed",
				  "code": 400
				}
			  }`,
			expectedOutput: nil,
			expectedError:  &xumm.ErrorResponse{ErrorResponseBody: xumm.ErrorResponseBody{Reference: "42d58b17-ee92-419d-b8ec-15797d10c4ed", Code: 400}},
			httpStatusCode: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {

			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.jsonResponse, tt.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			xapp := &Xapp{
				Cfg: cfg,
			}

			xp, err := xapp.PostXappPush(tt.request)

			if tt.expectedError != nil {
				assert.Nil(t, xp)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				body, _ := ioutil.ReadAll(m.Spy.Body)
				assert.JSONEq(t, tt.jsonRequest, string(body))
				assert.Equal(t, xp, tt.expectedOutput)
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
				assert.NoError(t, err)
			}
		})
	}
}