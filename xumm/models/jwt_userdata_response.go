package models

type JWTUserDataResponse struct {
	Operation string              `json:"operation,omitempty"`
	Data      JWTUserDataKeysData `json:"data,omitempty"`
	Keys      []string            `json:"keys,omitempty"`
	Count     int                 `json:"count,omitempty"`
	Persisted bool                `json:"persisted,omitempty"`
}

type JWTUserDataKeysData struct {
	Age  UserData `json:"age,omitempty"`
	Name UserData `json:"name,omitempty"`
}

type UserData struct {
	Value int    `json:"value,omitempty"`
	Name  string `json:"unit,omitempty"`
}
