package models

type AppStorageResponse struct {
	Application Application            `json:"application"`
	Stored      bool                   `json:"stored"`
	Data        map[string]interface{} `json:"data"`
}
