package models

type Application struct {
	UUIDV4     string `json:"uuidv4"`
	Name       string `json:"name"`
	WebhookUrl string `json:"webhookurl,omitempty"`
	Disabled   int    `json:"disabled,omitempty"`
}

type ApplicationDetails struct {
	Quota       map[string]interface{} `json:"quota,omitempty"`
	Application Application            `json:"application"`
	Call        Call                   `json:"call,omitempty"`
	JwtData     JwtPong                `json:"jwtData,omitempty"`
}

type AppDetails struct {
	UUIDv4     string  `json:"uuidv4"`
	Name       string  `json:"name"`
	WebhookURL *string `json:"webhookurl,omitempty"`
	Disabled   *int    `json:"disabled,omitempty"`
}
