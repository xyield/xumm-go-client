package models

import (
	anyjson "github.com/xyield/xumm-go-client/pkg/json"
)

type XummPayload struct {
	Meta        PayloadMeta        `json:"meta"`
	CustomMeta  XummCustomMeta     `json:"custom_meta"`
	Application PayloadApplication `json:"application"`
	Payload     Payload            `json:"payload"`
	Response    PayloadResponse    `json:"response"`
}

type XummDeletePayloadResponse struct {
	Meta       PayloadMeta      `json:"meta"`
	CustomMeta XummCustomMeta   `json:"custom_meta,omitempty"`
	Result     XummCancelResult `json:"result"`
}

type XummCancelResult struct {
	Cancelled bool   `json:"cancelled"`
	Reason    string `json:"reason"`
}

type PayloadMeta struct {
	Exists              bool        `json:"exists,omitempty"`
	UUID                string      `json:"uuid"`
	Multisign           bool        `json:"multisign,omitempty"`
	Submit              bool        `json:"submit,omitempty"`
	Destination         string      `json:"destination,omitempty"`
	ResolvedDestination string      `json:"resolved_destination,omitempty"`
	Finished            bool        `json:"finished,omitempty"`
	Resolved            bool        `json:"resolved,omitempty"`
	Signed              bool        `json:"signed,omitempty"`
	Cancelled           bool        `json:"cancelled,omitempty"`
	Expired             bool        `json:"expired,omitempty"`
	Pushed              bool        `json:"pushed,omitempty"`
	AppOpened           bool        `json:"app_opened,omitempty"`
	OpenedByDeeplink    interface{} `json:"opened_by_deeplink,omitempty"`
	ReturnURLApp        string      `json:"return_url_app,omitempty"`
	ReturnURLWeb        interface{} `json:"return_url_web,omitempty"`
	CustomIdentifier    string      `json:"custom_identifier,omitempty"`
	CustomBlob          string      `json:"custom_blob,omitempty"`
	CustomInstruction   string      `json:"custom_instruction,omitempty"`
	IsXapp              bool        `json:"is_xapp,omitempty"`
}

type PayloadApplication struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Disabled        int    `json:"disabled"`
	Uuidv4          string `json:"uuidv4"`
	IconURL         string `json:"icon_url"`
	IssuedUserToken string `json:"issued_user_token"`
}

type Payload struct {
	TxType           string          `json:"tx_type"`
	TxDestination    string          `json:"tx_destination"`
	TxDestinationTag int             `json:"tx_destination_tag"`
	RequestJSON      anyjson.AnyJson `json:"request_json"`
	Origintype       string          `json:"origintype"`
	Signmethod       string          `json:"signmethod"`
	CreatedAt        string          `json:"created_at"`
	ExpiresAt        string          `json:"expires_at"`
	ExpiresInSeconds int             `json:"expires_in_seconds"`
}

type PayloadResponse struct {
	Hex                string `json:"hex"`
	Txid               string `json:"txid"`
	ResolvedAt         string `json:"resolved_at"`
	DispatchedTo       string `json:"dispatched_to"`
	DispatchedNodetype string `json:"dispatched_nodetype"`
	DispatchedResult   string `json:"dispatched_result"`
	MultisignAccount   string `json:"multisign_account"`
	Account            string `json:"account"`
}
