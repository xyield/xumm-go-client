package xumm

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xumm-go-client/models"
)

func TestXrplTx(t *testing.T) {
	os.Setenv("XUMM_API_KEY", "testApiKey")
	os.Setenv("XUMM_API_SECRET", "testApiSecret")

	bc := &models.BalanceDetails{
		Value:        "-1.000012",
		Currency:     "XRP",
		CounterParty: "",
	}

	txJson := &models.Transaction{
		"Account":         "r4bA4uZgXadPMzURqGLCvCmD48FmXJWHCG",
		"Amount":          "1000000",
		"Destination":     "rPdvC6ccq8hCdPKSPJkPmyZ4Mi1oG2FFkT",
		"Fee":             "12",
		"Flags":           int64(2147483648),
		"Sequence":        int64(58549314),
		"SigningPubKey":   "0260F06C0590C470E7E7FA9DE3D9E85B1825E19196D8893DD84431F6E9491739AC",
		"TransactionType": "Payment",
		"meta": map[string]interface{}{
			"TransactionIndex":  int64(0),
			"TransactionResult": "tesSUCCESS",
			"delivered_amount":  "1000000",
		},
		"validated": true,
	}

	json := `{
		"txid": "A17E4DEAD62BF705D9B73B4EAD2832F1C55C6C5A0067327A45E497FD8D31C0E3",
		"balanceChanges": {
		  "r4bA4uZgXadPMzURqGLCvCmD48FmXJWHCG": [
			{
			  "counterparty": "",
			  "currency": "XRP",
			  "value": "-1.000012"
			}
		  ]
		},
		"node": "wss://xrpl.ws",
		"transaction": {
		  "Account": "r4bA4uZgXadPMzURqGLCvCmD48FmXJWHCG",
		  "Amount": "1000000",
		  "Destination": "rPdvC6ccq8hCdPKSPJkPmyZ4Mi1oG2FFkT",
		  "Fee": "12",
		  "Flags": 2147483648,
		  "Sequence": 58549314,
		  "SigningPubKey": "0260F06C0590C470E7E7FA9DE3D9E85B1825E19196D8893DD84431F6E9491739AC",
		  "TransactionType": "Payment",
		  "meta": {
			"TransactionIndex": 0,
			"TransactionResult": "tesSUCCESS",
			"delivered_amount": "1000000"
		  },
		  "validated": true
		}
	  }`

	txRes := &models.XrpTxResponse{
		Txid: "A17E4DEAD62BF705D9B73B4EAD2832F1C55C6C5A0067327A45E497FD8D31C0E3",
		Node: "wss://xrpl.ws",
		BalanceChanges: map[string][]models.BalanceDetails{
			"r4bA4uZgXadPMzURqGLCvCmD48FmXJWHCG": {
				*bc,
			},
		},
		Transaction: *txJson,
	}

	var tests = []struct {
		testName       string
		input          string
		json           string
		expectedOutput *models.XrpTxResponse
		expectedError  error
		httpStatusCode int
	}{
		{testName: "valid transaction id", input: "A17E4DEAD62BF705D9B73B4EAD2832F1C55C6C5A0067327A45E497FD8D31C0E3", json: json, expectedOutput: txRes, expectedError: nil, httpStatusCode: 200},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			m := &MockClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					b := ioutil.NopCloser(bytes.NewBuffer([]byte(tt.json)))
					return &http.Response{
						StatusCode: tt.httpStatusCode,
						Body:       b,
					}, nil
				},
			}
			c, _ := NewClient(WithHttpClient(m))

			tx, err := c.XrplTransaction(tt.input)

			if tt.expectedError != nil {
				assert.Nil(t, tx)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, tx)
			}

		})
	}
}
