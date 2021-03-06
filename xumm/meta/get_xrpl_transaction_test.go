//go:build unit
// +build unit

package meta

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	testutils "github.com/xyield/xumm-go-client/utils/test-utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestGetXrplTx(t *testing.T) {

	txJson := &anyjson.AnyJson{
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

	f := models.Formatted{
		Value:    "-0.589015",
		Currency: "XRP",
	}
	bc := &models.BalanceDetails{
		Value:        "-1.000012",
		Currency:     "XRP",
		CounterParty: "",
		Formatted:    f,
	}
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

	json := testutils.ConvertJsonFileToJsonString("static-test-data/xrpl_transaction_test.json")

	errorJson := `{
		"error": {
		  "reference": "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c",
		  "code": 812
		}
	  }`

	var tests = []struct {
		testName       string
		input          string
		json           string
		expectedOutput *models.XrpTxResponse
		expectedError  error
		httpStatusCode int
	}{
		{testName: "valid transaction id", input: "A17E4DEAD62BF705D9B73B4EAD2832F1C55C6C5A0067327A45E497FD8D31C0E3", json: json, expectedOutput: txRes, expectedError: nil, httpStatusCode: 200},
		{testName: "empty transaction id", input: "", json: "", expectedOutput: nil, expectedError: &EmptyTransactionId{}, httpStatusCode: 0},
		{testName: "error response", input: "A17E4DEAD62BF705D9B73B4EAD2832F1C55C6C5A0067327A45E497FD8D31C0E3", json: errorJson, expectedOutput: nil, expectedError: &xumm.ErrorResponse{ErrorResponseBody: xumm.ErrorResponseBody{Reference: "3a04c7d3-94aa-4d8d-9559-62bb5e8a653c", Code: 812}}, httpStatusCode: 403},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			m := &testutils.MockClient{}
			m.DoFunc = testutils.MockResponse(tt.json, tt.httpStatusCode, m)
			cfg, err := xumm.NewConfig(xumm.WithHttpClient(m), xumm.WithAuth("testApiKey", "testApiSecret"))
			assert.NoError(t, err)
			meta := &Meta{
				Cfg: cfg,
			}
			tx, err := meta.GetXrplTransaction(tt.input)

			if tt.expectedError != nil {
				assert.Nil(t, tx)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedOutput, tx)
				assert.Equal(t, http.Header{
					"X-API-Key":    {"testApiKey"},
					"X-API-Secret": {"testApiSecret"},
					"Content-Type": {"application/json"},
				}, m.Spy.Header)
			}

		})
	}
}
