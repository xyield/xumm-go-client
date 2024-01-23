package models

type Pong struct {
	Pong bool               `json:"pong"`
	Auth ApplicationDetails `json:"auth"`
}

type Call struct {
	UUIDV4 string `json:"uuidv4"`
}
