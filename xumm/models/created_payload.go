package models

type CreatedPayload struct {
	UUID   string `json:"uuid"`
	Next   Next   `json:"next"`
	Refs   Refs   `json:"refs"`
	Pushed bool   `json:"pushed"`
}

type Next struct {
	Always            string `json:"always"`
	NoPushMsgReceived string `json:"no_push_msg_received"`
}

type Refs struct {
	QrPng            string   `json:"qr_png"`
	QrMatrix         string   `json:"qr_matrix"`
	QrURIQualityOpts []string `json:"qr_uri_quality_opts"`
	WebsocketStatus  string   `json:"websocket_status"`
}
