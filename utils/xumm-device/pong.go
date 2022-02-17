package xummdevice

type Pong struct {
	Pong  bool `json:"pong"`
	Badge int  `json:"badge"`
	Auth  struct {
		User struct {
			Uuidv4 string `json:"uuidv4"`
			Slug   string `json:"slug"`
			Name   string `json:"name"`
		} `json:"user"`
		Device struct {
			Uuidv4      string `json:"uuidv4"`
			Idempotence int    `json:"idempotence"`
		} `json:"device"`
		Call struct {
			Hash        string `json:"hash"`
			Idempotence int    `json:"idempotence"`
			Uuidv4      string `json:"uuidv4"`
		} `json:"call"`
	} `json:"auth"`
}
