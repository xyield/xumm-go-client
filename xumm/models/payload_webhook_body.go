package models

import anyjson "github.com/xyield/xumm-go-client/utils/json"

type PayloadWebhookBody struct {
	Meta            WebHookMeta            `json:"meta"`
	CustomMeta      CustomMeta             `json:"custom_meta"`
	PayloadResponse WebHookPayloadResponse `json:"payloadResponse"`
	UserToken       UserToken              `json:"userToken"`
}

type WebHookMeta struct {
	URL               string `json:"url"`
	ApplicationUuidv4 string `json:"application_uuidv4"`
	PayloadUuidv4     string `json:"payload_uuidv4"`
	OpenedByDeeplink  bool   `json:"opened_by_deeplink,omitempty"`
}

type CustomMeta struct {
	Identifier  string          `json:"identifier,omitempty"`
	Blob        anyjson.AnyJson `json:"blob,omitempty"`
	Instruction string          `json:"instruction,omitempty"`
}

type WebHookPayloadResponse struct {
	PayloadUuidv4       string    `json:"payload_uuidv4"`
	ReferenceCallUuidv4 string    `json:"reference_call_uuidv4"`
	Signed              bool      `json:"signed"`
	UserToken           bool      `json:"user_token"`
	ReturnURL           ReturnURL `json:"return_url"`
	Txid                string    `json:"txid"`
}

type ReturnURL struct {
	App string `json:"app,omitempty"`
	Web string `json:"web,omitempty"`
}

type UserToken struct {
	UserToken       string `json:"user_token"`
	TokenIssued     int    `json:"token_issued"`
	TokenExpiration int    `json:"token_expiration,omitempty"`
}
