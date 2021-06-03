package models

type Pong struct {
	Pong bool               `json:"pong"`
	Auth ApplicationDetails `json:"auth"`
}

type ApplicationDetails struct {
	Quota       map[string]interface{} `json:"quota"`
	Application Application            `json:"application"`
	Call        Call                   `json:"call"`
}

type Application struct {
	UUIDV4     string `json:"uuidv4"`
	Name       string `json:"name"`
	WebhookUrl string `json:"webhookurl"`
	Disabled   int
}

type Call struct {
	UUIDV4 string `json:"uuidv4"`
}
