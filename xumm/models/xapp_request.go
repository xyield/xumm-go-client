package models

import anyjson "github.com/xyield/xumm-go-client/utils/json"

type XappRequest struct {
	UserToken string          `json:"user_token"`
	Subtitle  string          `json:"subtitle"`
	Body      string          `json:"body,omitempty"`
	Data      anyjson.AnyJson `json:"data,omitempty"`
	Silent    *bool           `json:"silent,omitempty"`
}
