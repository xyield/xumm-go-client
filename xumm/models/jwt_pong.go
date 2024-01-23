package models

type JwtPong struct {
	Pong            bool   `json:"pong"`
	ClientID        string `json:"client_id"`
	State           string `json:"state,omitempty"`
	Scope           string `json:"scope,omitempty"`
	Nonce           string `json:"nonce,omitempty"`
	Aud             string `json:"aud"`
	Sub             string `json:"sub"`
	AppUUIDv4       string `json:"app_uuidv4"`
	AppName         string `json:"app_name"`
	PayloadUUIDv4   string `json:"payload_uuidv4"`
	UserTokenUUIDv4 string `json:"usertoken_uuidv4"`
	Iat             int    `json:"iat"`
	Exp             int    `json:"exp"`
	Iss             string `json:"iss"`
}
