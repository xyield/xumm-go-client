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

func TestGetHookHashes(t *testing.T) {

	noHookHash := `{
		"31C3EC186C367DA66DFBD0E576D6170A2C1AB846BFC35FC0B49D202F2A8CDFD8":	{
			"name":	"Savings",
			"description":	"Efficiently manage your finances with the savings hook by automatically transferring a percentage of your transaction amount to a separate savings account to prevent accidental spending",
			"creator":	{
				"name":	"XRPL Labs",
				"mail":	"support@xrpl-labs.com",
				"site":	"xrpl-labs.com"
			},
			"xapp":	"hooks.savings",
			"appuuid":	"d21fe83d-64b9-4aef-8908-9a4952d8922c",
			"icon":	"https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/1857a4e1-6711-4c37-bb54-7664e129e9bf.png",
			"verifiedAccounts":	[],
			"audits":	[]
		},
		"FCEA17EB0A771899227CD1C22CD1BFC9C9B78017E6EAE0F576C6E78BAE9D57D4":	{
			"name":	"Firewall",
			"description":	"Select which type of transactions you allow into and out of your account.",
			"creator":	{
				"name":	"XRPL Labs",
				"mail":	"support@xrpl-labs.com",
				"site":	"xrpl-labs.com"
			},
			"xapp":	"hooks.firewall",
			"appuuid":	"26447750-0063-4f81-9de4-c7a9bb8fe935",
			"icon":	"https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/9a1800d8-d2a3-4bbb-9cdd-e8e47850d65c.png",
			"verifiedAccounts":	[],
			"audits":	[]
		},
		"C08FF0A3849F7D4A2F346DDB69282487BA320557516DC954C91CA81F05F478BB":	{
			"name":	"Direct Debit",
			"description":	"Allow trusted third parties to pull funds from your account up to a limit you set. For example your power company can bill you and your account can automatically pay that bill.",
			"creator":	{
				"name":	"XRPL Labs",
				"mail":	"support@xrpl-labs.com",
				"site":	"xrpl-labs.com"
			},
			"xapp":	"hooks.directdebit",
			"appuuid":	"6df15d3e-7c6e-49a5-aceb-63df92e92fc2",
			"icon":	"https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/24738ef6-df4b-464c-9a07-5c9f01c76f00.png",
			"verifiedAccounts":	[],
			"audits":	[]
		},
		"9B11CAAF4F87D2638CE86D3E82FE4A1464D95693B4C440B9048DCC5EFD91862B":	{
			"name":	"Payment Watchdog",
			"description":	"When sending high value transactions out of your account, require first a notification that a high valued payment will be made, followed by a time delay, followed by the high value transaction itself. This prevents accidental high value sends, adding an additional layer of security to your account.",
			"creator":	{
				"name":	"XRPL Labs",
				"mail":	"support@xrpl-labs.com",
				"site":	"xrpl-labs.com"
			},
			"xapp":	"hooks.highvalueprotect",
			"appuuid":	"dcfe584f-8158-49ec-840a-3aea8819ce72",
			"icon":	"https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/1857a4e1-6711-4c37-bb54-7664e129e9bf.png",
			"verifiedAccounts":	[],
			"audits":	[]
		},
		"???":	{
			"name":	"Balance Adjustment",
			"description":	"Claim Xahau Balance Adjustment.",
			"creator":	{
				"name":	"Xahau",
				"mail":	"mail@xahau.network",
				"site":	"xahau.network"
			},
			"xapp":	"xahau.balanceadjustment",
			"appuuid":	"3b58aa5d-5263-4252-b230-f81ab0db9a7a",
			"icon":	"https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/3b1fd165-1923-4ed2-a48c-80709cd6a127.png",
			"verifiedAccounts":	[],
			"audits":	[]
		}
	}`

	hookhashresp := models.HookHashesResponse{
		"31C3EC186C367DA66DFBD0E576D6170A2C1AB846BFC35FC0B49D202F2A8CDFD8": models.HookHashResponse{
			Name:        "Savings",
			Description: "Efficiently manage your finances with the savings hook by automatically transferring a percentage of your transaction amount to a separate savings account to prevent accidental spending",
			Creator: models.HookHashCreator{
				Name: "XRPL Labs",
				Mail: "support@xrpl-labs.com",
				Site: "xrpl-labs.com",
			},
			Xapp:             "hooks.savings",
			AppUUID:          "d21fe83d-64b9-4aef-8908-9a4952d8922c",
			Icon:             "https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/1857a4e1-6711-4c37-bb54-7664e129e9bf.png",
			VerifiedAccounts: []string{},
			Audits:           []string{},
		},
		"FCEA17EB0A771899227CD1C22CD1BFC9C9B78017E6EAE0F576C6E78BAE9D57D4": models.HookHashResponse{
			Name:        "Firewall",
			Description: "Select which type of transactions you allow into and out of your account.",
			Creator: models.HookHashCreator{
				Name: "XRPL Labs",
				Mail: "support@xrpl-labs.com",
				Site: "xrpl-labs.com",
			},
			Xapp:             "hooks.firewall",
			AppUUID:          "26447750-0063-4f81-9de4-c7a9bb8fe935",
			Icon:             "https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/9a1800d8-d2a3-4bbb-9cdd-e8e47850d65c.png",
			VerifiedAccounts: []string{},
			Audits:           []string{},
		},
		"C08FF0A3849F7D4A2F346DDB69282487BA320557516DC954C91CA81F05F478BB": models.HookHashResponse{
			Name:        "Direct Debit",
			Description: "Allow trusted third parties to pull funds from your account up to a limit you set. For example your power company can bill you and your account can automatically pay that bill.",
			Creator: models.HookHashCreator{
				Name: "XRPL Labs",
				Mail: "support@xrpl-labs.com",
				Site: "xrpl-labs.com",
			},
			Xapp:             "hooks.directdebit",
			AppUUID:          "6df15d3e-7c6e-49a5-aceb-63df92e92fc2",
			Icon:             "https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/24738ef6-df4b-464c-9a07-5c9f01c76f00.png",
			VerifiedAccounts: []string{},
			Audits:           []string{},
		},
		"9B11CAAF4F87D2638CE86D3E82FE4A1464D95693B4C440B9048DCC5EFD91862B": models.HookHashResponse{
			Name:        "Payment Watchdog",
			Description: "When sending high value transactions out of your account, require first a notification that a high valued payment will be made, followed by a time delay, followed by the high value transaction itself. This prevents accidental high value sends, adding an additional layer of security to your account.",
			Creator: models.HookHashCreator{
				Name: "XRPL Labs",
				Mail: "support@xrpl-labs.com",
				Site: "xrpl-labs.com",
			},
			Xapp:             "hooks.highvalueprotect",
			AppUUID:          "dcfe584f-8158-49ec-840a-3aea8819ce72",
			Icon:             "https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/1857a4e1-6711-4c37-bb54-7664e129e9bf.png",
			VerifiedAccounts: []string{},
			Audits:           []string{},
		},
		"???": models.HookHashResponse{
			Name:        "Balance Adjustment",
			Description: "Claim Xahau Balance Adjustment.",
			Creator: models.HookHashCreator{
				Name: "Xahau",
				Mail: "mail@xahau.network",
				Site: "xahau.network",
			},
			Xapp:             "xahau.balanceadjustment",
			AppUUID:          "3b58aa5d-5263-4252-b230-f81ab0db9a7a",
			Icon:             "https://cdn.xumm.pro/cdn-cgi/image/width=500,height=500,quality=75,fit=crop/app-logo/3b1fd165-1923-4ed2-a48c-80709cd6a127.png",
			VerifiedAccounts: []string{},
			Audits:           []string{},
		},
	}

	var tests = []struct {
		testName       string
		inputValue     string
		expectedOutput *models.HookHashesResponse
		expectedError  error
		httpStatusCode int
	}{
		{
			testName:       "GetHookHashes - success",
			inputValue:     noHookHash,
			expectedOutput: &hookhashresp,
			expectedError:  nil,
			httpStatusCode: http.StatusOK,
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
			ca, err := meta.GetHookHashes()

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
