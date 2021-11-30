package models

type XappEventResponse struct {
	Pushed bool   `json:"pushed"`
	UUID   string `json:"uuid"`
}
