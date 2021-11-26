package payload

import (
	"testing"

	anyjson "github.com/xyield/xumm-go-client/pkg/json"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func TestDeletePayloadByUuid(t *testing.T) {
	tt := []struct {
		description    string
		uuid           string
		jsonResponse   string
		expectedOutput *models.XummDeletePayloadResponse
		statusCode     int
		expectedError  error
	}{
		{
			description: "successfully cancelled",
			uuid:        "XXX",
			jsonResponse: `{
				"result": {
				  "cancelled": true,
				  "reason": "OK"
				},
				"meta": {
				  "exists": true,
				  "uuid": "<some-uuid>",
				  "multisign": false,
				  "submit": true,
				  "destination": "rPEPPER7kfTD9w2To4CQk6UCfuHM9c6GDY",
				  "resolved_destination": "XRP Tip Bot",
				  "finished": false,
				  "expired": true,
				  "pushed": true,
				  "app_opened": false,
				  "return_url_app": "<some-url-or-null>",
				  "return_url_web": "<some-url-or-null>"
				},
				"custom_meta": {
				  "identifier": "some_identifier_1337",
				  "blob": {},
				  "instruction": "Hey ❤️ ..."
				}
			  }`,
			expectedOutput: &models.XummDeletePayloadResponse{
				Result: models.XummCancelResult{
					Cancelled: true,
					Reason:    "OK",
				},
				Meta: models.PayloadMeta{
					Exists:              true,
					UUID:                "<some-uuid>",
					Multisign:           false,
					Submit:              true,
					Destination:         "rPEPPER7kfTD9w2To4CQk6UCfuHM9c6GDY",
					ResolvedDestination: "XRP Tip Bot",
					Finished:            false,
					Expired:             true,
					Pushed:              true,
					AppOpened:           false,
					ReturnURLApp:        "<some-url-or-null>",
					ReturnURLWeb:        "<some-url-or-null>",
				},
				CustomMeta: models.XummCustomMeta{
					Identifier:  "some_identifier_1337",
					Blob:        anyjson.AnyJson{},
					Instruction: "Hey ❤️ ...",
				},
			},
			statusCode:    200,
			expectedError: nil,
		},
		{
			description: "OK but not cancelled",
			uuid:        "XXX",
			jsonResponse: `{
				"result": {
				  "cancelled": false,
				  "reason": "<some-reason-see-note-below>"
				},
				"meta": {
				  "exists": true,
				  "uuid": "<some-uuid>",
				  "multisign": false,
				  "submit": true,
				  "destination": "rPEPPER7kfTD9w2To4CQk6UCfuHM9c6GDY",
				  "resolved_destination": "XRP Tip Bot",
				  "finished": false,
				  "expired": true,
				  "pushed": true,
				  "app_opened": false,
				  "return_url_app": "<some-url-or-null>",
				  "return_url_web": "<some-url-or-null>",
				  "custom_identifier": "some_identifier_1337",
				  "custom_blob": {},
				  "custom_instruction": "Hey ❤️ ..."
				}
			  }`,
			expectedOutput: &models.XummDeletePayloadResponse{
				Result: models.XummCancelResult{
					Cancelled: false,
					Reason:    "<some-reason-see-note-below>",
				},
				Meta: models.PayloadMeta{
					Exists:              true,
					UUID:                "<some-uuid>",
					Multisign:           false,
					Submit:              true,
					Destination:         "rPEPPER7kfTD9w2To4CQk6UCfuHM9c6GDY",
					ResolvedDestination: "XRP Tip Bot",
					Finished:            false,
					Expired:             true,
					Pushed:              true,
					AppOpened:           false,
					ReturnURLApp:        "<some-url-or-null>",
					ReturnURLWeb:        "<some-url-or-null>",
					CustomIdentifier:    "some_identifier_1337",
					CustomBlob:          anyjson.AnyJson{},
					CustomInstruction:   "Hey ❤️ ...",
				},
			},
			statusCode:    200,
			expectedError: nil,
		},
	}
}
