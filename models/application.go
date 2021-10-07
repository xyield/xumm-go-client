package models

type Application struct {
	UUIDV4     string `json:"uuidv4"`
	Name       string `json:"name"`
	WebhookUrl string `json:"webhookurl,omitempty"`
	Disabled   int    `json:"disabled,omitempty"`
}
