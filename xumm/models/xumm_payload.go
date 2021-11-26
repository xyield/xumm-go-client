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

type PayloadMeta struct {
	Exists              bool        `json:"exists"`
	UUID                string      `json:"uuid"`
	Multisign           bool        `json:"multisign"`
	Submit              bool        `json:"submit"`
	Destination         string      `json:"destination"`
	ResolvedDestination string      `json:"resolved_destination"`
	Resolved            bool        `json:"resolved"`
	Signed              bool        `json:"signed"`
	Cancelled           bool        `json:"cancelled"`
	Expired             bool        `json:"expired"`
	Pushed              bool        `json:"pushed"`
	AppOpened           bool        `json:"app_opened"`
	OpenedByDeeplink    interface{} `json:"opened_by_deeplink"`
	ReturnURLApp        string      `json:"return_url_app"`
	ReturnURLWeb        interface{} `json:"return_url_web"`
	IsXapp              bool        `json:"is_xapp"`
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
