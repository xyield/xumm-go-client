package models

type UserTokenValidity struct {
	UserToken string `json:"user_token"`
	Active    bool   `json:"active"`
	Issued    int64  `json:"issued,omitempty"`
	Expires   int64  `json:"expires,omitempty"`
}

type UserTokenResponse struct {
	Tokens []UserTokenValidity `json:"tokens"`
}

type UserTokenRequest struct {
	Tokens []string `json:"tokens"`
}
