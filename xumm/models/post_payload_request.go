package models

import anyjson "github.com/xyield/xumm-go-client/utils/json"

type XummPostPayload struct {
	UserToken  string                  `json:"user_token,omitempty"`
	TxJson     anyjson.AnyJson         `json:"txjson,omitempty"`
	TxBlob     string                  `json:"txblob,omitempty"`
	Options    *XummPostPayloadOptions `json:"options,omitempty"`
	CustomMeta *XummCustomMeta         `json:"custom_meta,omitempty"`
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
	Identifier  string          `json:"identifier"`
	Blob        anyjson.AnyJson `json:"blob"`
	Instruction string          `json:"instruction"`
}
