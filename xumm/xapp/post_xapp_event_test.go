package xapp

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestPostXappEvent(t *testing.T) {

	f := new(bool)
	*f = false

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
			description: "successfully create event with 'silent' set to false",
			request: models.XappRequest{
				UserToken: "token",
				Subtitle:  "subtitle",
				Body:      "body",
				Data:      anyjson.AnyJson{},
				Silent:    f,
			},
			jsonRequest: `{
				"user_token": "token",
				"subtitle": "subtitle",
				"body": "body",
				"silent": false
			}`,
			jsonResponse: `{
				"pushed": true,
				"uuid": "token"
			  }`,
			expectedOutput: &models.XappResponse{Pushed: true, UUID: "token"},
			expectedError:  nil,
			httpStatusCode: 200,
		},
		{
			description: "successfully create event without 'silent' set",
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
				"pushed": true,
				"uuid": "token"
			  }`,
			expectedOutput: &models.XappResponse{Pushed: true, UUID: "token"},
			expectedError:  nil,
			httpStatusCode: 200,
		},
		{
			description: "invalid request body",
			request: models.XappRequest{
				UserToken: "",
				Subtitle:  "",
				Body:      "body",
				Data:      anyjson.AnyJson{},
				Silent:    f,
			},
			jsonRequest:    "",
			jsonResponse:   "",
			expectedOutput: nil,
			expectedError:  &invalidEventRequestError{},
			httpStatusCode: 200,
		},
		{
			description: "error creating event",
			request: models.XappRequest{
				UserToken: "token",
				Subtitle:  "subtitle",
				Body:      "body",
				Data:      anyjson.AnyJson{},
				Silent:    f,
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

			xe, err := xapp.PostXappEvent(tt.request)

			if tt.expectedError != nil {
				assert.Nil(t, xe)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				body, _ := ioutil.ReadAll(m.Spy.Body)
				assert.JSONEq(t, tt.jsonRequest, string(body))
				assert.Equal(t, xe, tt.expectedOutput)
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
