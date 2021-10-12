package xumm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
)

func TestGetAppStorage(t *testing.T) {
	tests := []struct {
		description    string
		response       string
		expectedOutput *models.AppStorageResponse
		expectedError  error
	}{
		{
			description: "Return app storage with no data",
			response: `{
				"application": {
				  "name": "My App",
				  "uuidv4": "8525e32b-xxxx-xxxx-xxxx-f794874a80b0"
				},
				"stored": true,
				"data": null
			  }`,
			expectedOutput: &models.AppStorageResponse{
				Application: models.Application{
					Name:   "My App",
					UUIDV4: "8525e32b-xxxx-xxxx-xxxx-f794874a80b0",
				},
				Stored: true,
				Data:   nil,
			},
			expectedError: nil,
		},
		{
			description: "Return app storage with empty data object",
			response: `{
				"application": {
				  "name": "My App",
				  "uuidv4": "8525e32b-xxxx-xxxx-xxxx-f794874a80b0"
				},
				"stored": true,
				"data": {}
			  }`,
			expectedOutput: &models.AppStorageResponse{
				Application: models.Application{
					Name:   "My App",
					UUIDV4: "8525e32b-xxxx-xxxx-xxxx-f794874a80b0",
				},
				Stored: true,
				Data:   map[string]interface{}{},
			},
			expectedError: nil,
		},
		{
			description: "Return app storage with data object with one key value pair",
			response: `{
				"application": {
				  "name": "My App",
				  "uuidv4": "8525e32b-xxxx-xxxx-xxxx-f794874a80b0"
				},
				"stored": true,
				"data": {
					"someKey": "someValue"
				}
			  }`,
			expectedOutput: &models.AppStorageResponse{
				Application: models.Application{
					Name:   "My App",
					UUIDV4: "8525e32b-xxxx-xxxx-xxxx-f794874a80b0",
				},
				Stored: true,
				Data: map[string]interface{}{
					"someKey": "someValue",
				},
			},
			expectedError: nil,
		},
		{
			description: "Return app storage with data object with multiple different data types",
			response: `{
				"application": {
				  "name": "My App",
				  "uuidv4": "8525e32b-xxxx-xxxx-xxxx-f794874a80b0"
				},
				"stored": true,
				"data": {
					"someKey": "someValue",
					"float": 10.500,
					"int": 1337
				}
			  }`,
			expectedOutput: &models.AppStorageResponse{
				Application: models.Application{
					Name:   "My App",
					UUIDV4: "8525e32b-xxxx-xxxx-xxxx-f794874a80b0",
				},
				Stored: true,
				Data: map[string]interface{}{
					"someKey": "someValue",
					"int":     int64(1337),
					"float":   10.500,
				},
			},
			expectedError: nil,
		},
		{
			description: "Return app storage with data object with nested objects in data",
			response: `{
				"application": {
				  "name": "My App",
				  "uuidv4": "8525e32b-xxxx-xxxx-xxxx-f794874a80b0"
				},
				"stored": true,
				"data": {
					"someKey": "someValue",
					"float": 10.500,
					"nested": {
						"someKey": "someValue",
						"float": 12.0,
						"int": 9182
					}
				}
			  }`,
			expectedOutput: &models.AppStorageResponse{
				Application: models.Application{
					Name:   "My App",
					UUIDV4: "8525e32b-xxxx-xxxx-xxxx-f794874a80b0",
				},
				Stored: true,
				Data: map[string]interface{}{
					"someKey": "someValue",
					"float":   10.500,
					"nested": map[string]interface{}{
						"someKey": "someValue",
						"float":   12.0,
						"int":     int64(9182),
					},
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			mockClient := &MockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					b := ioutil.NopCloser(bytes.NewReader([]byte(tt.response)))
					return &http.Response{
						StatusCode: 200,
						Body:       b,
					}, nil
				},
			}
			c, _ := NewClient(WithHttpClient(mockClient), WithAuth("testApiKey", "testApiSecret"))
			as, err := c.GetAppStorage()

			if tt.expectedError != nil {
				assert.Nil(t, as)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, as)
			}
		})
	}
}

func TestSetAppStorage(t *testing.T) {
	tests := []struct {
		description    string
		input          map[string]interface{}
		response       string
		expectedOutput *models.AppStorageResponse
		expectedError  error
	}{
		{
			description: "Set single key value pair with string data type",
			input: map[string]interface{}{
				"someKey": "someValue",
			},
			response: `{
				"application": {
				  "name": "My App",
				  "uuidv4": "8525e32b-xxxx-xxxx-xxxx-f794874a80b0"
				},
				"stored": true,
				"data": {
					"someKey": "someValue"
				}
			}`,
			expectedOutput: &models.AppStorageResponse{
				Application: models.Application{
					Name:   "My App",
					UUIDV4: "8525e32b-xxxx-xxxx-xxxx-f794874a80b0",
				},
				Stored: true,
				Data: map[string]interface{}{
					"someKey": "someValue",
				},
			},
			expectedError: nil,
		},
		{
			description: "Set multiple key value pair with various data type",
			input: map[string]interface{}{
				"someKey":   "someValue",
				"someInt":   int64(1337),
				"someFloat": 13.07,
			},
			response: `{
				"application": {
				  "name": "My App",
				  "uuidv4": "8525e32b-xxxx-xxxx-xxxx-f794874a80b0"
				},
				"stored": true,
				"data": {
					"someKey": "someValue",
					"someFloat": 13.07,
					"someInt": 1337
				}
			}`,
			expectedOutput: &models.AppStorageResponse{
				Application: models.Application{
					Name:   "My App",
					UUIDV4: "8525e32b-xxxx-xxxx-xxxx-f794874a80b0",
				},
				Stored: true,
				Data: map[string]interface{}{
					"someKey":   "someValue",
					"someFloat": 13.07,
					"someInt":   int64(1337),
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			mockClient := &MockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					b := ioutil.NopCloser(bytes.NewReader([]byte(tt.response)))
					return &http.Response{
						StatusCode: 200,
						Body:       b,
					}, nil
				},
			}
			c, _ := NewClient(WithHttpClient(mockClient), WithAuth("testApiKey", "testApiSecret"))
			as, err := c.SetAppStorage(tt.input)

			if tt.expectedError != nil {
				assert.Nil(t, as)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, as)
			}
		})
	}
}

func TestDeleteAppStorage(t *testing.T) {
	tests := []struct {
		description    string
		response       string
		expectedOutput *models.AppStorageResponse
		expectedError  error
	}{
		{
			description: "Delete data from app storage",
			response: `{
				"application": {
				  "name": "My App",
				  "uuidv4": "8525e32b-xxxx-xxxx-xxxx-f794874a80b0"
				},
				"stored": true,
				"data": null
			}`,
			expectedOutput: &models.AppStorageResponse{
				Application: models.Application{
					Name:   "My App",
					UUIDV4: "8525e32b-xxxx-xxxx-xxxx-f794874a80b0",
				},
				Stored: true,
				Data:   nil,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			mockClient := &MockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					b := ioutil.NopCloser(bytes.NewReader([]byte(tt.response)))
					return &http.Response{
						StatusCode: 200,
						Body:       b,
					}, nil
				},
			}
			c, _ := NewClient(WithHttpClient(mockClient), WithAuth("testApiKey", "testApiSecret"))
			as, err := c.DeleteAppStorage()

			if tt.expectedError != nil {
				assert.Nil(t, as)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, as)
			}
		})
	}
}
