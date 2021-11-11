package models

import "github.com/xyield/xumm-go-client/pkg/json"

type XummPostPayload struct {
	UserToken  string                  `json:"user_token,omitempty"`
	TxJson     json.AnyJson            `json:"txjson,omitempty"`
	TxBlob     string                  `json:"txblob,omitempty"`
	Options    *XummPostPayloadOptions `json:"options,omitempty"`
	CustomMeta *XummCustomMeta         `json:"custom_meta,omitempty"`
}

type JsonTransaction struct {
	TransactionType string `json:"TransactionType,omitempty"`
}

type XummPostPayloadOptions struct {
	Submit    bool                             `json:"submit,omitempty"`
	Multisign bool                             `json:"multisign,omitempty"`
	Expire    int32                            `json:"expire,omitempty"`
	ReturnUrl *XummPostPayloadOptionsReturnUrl `json:"return_url,omitempty"`
}

type XummPostPayloadOptionsReturnUrl struct {
	App string `json:"app,omitempty"`
	Web string `json:"web,omitempty"`
}

type XummCustomMeta struct {
	Identifier  string       `json:"identifier"`
	Blob        json.AnyJson `json:"blob"`
	Instruction string       `json:"instruction"`
}
