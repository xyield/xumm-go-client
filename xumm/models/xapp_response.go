package models

type XappResponse struct {
	Pushed bool   `json:"pushed"`
	UUID   string `json:"uuid,omitempty"`
}
