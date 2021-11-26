package models

import anyjson "github.com/xyield/xumm-go-client/pkg/json"

type AppStorageResponse struct {
	Application Application     `json:"application"`
	Stored      bool            `json:"stored"`
	Data        anyjson.AnyJson `json:"data"`
}
