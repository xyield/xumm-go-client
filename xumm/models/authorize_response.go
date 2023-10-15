package models

type XappJWTAuthorizeResponse struct {
	OTT OTT                         `json:"ott,omitempty"`
	App XappJWTAuthorizeResponseApp `json:"app,omitempty"`
	JWT string                      `json:"jwt,omitempty"`
}

type OTT struct {
	Account       string                         `json:"account,omitempty"`
	AccountAccess string                         `json:"account_access,omitempty"`
	AccountType   string                         `json:"account_type,omitempty"`
	Locale        string                         `json:"locale,omitempty"`
	Style         string                         `json:"style,omitempty"`
	Version       string                         `json:"version,omitempty"`
	Origin        XappJWTAuthorizeResponseOrigin `json:"origin,omitempty"`
	User          string                         `json:"user,omitempty"`
	UserDevice    UserDevice                     `json:"user_device,omitempty"`
}

type XappJWTAuthorizeResponseOrigin struct {
	Type string                       `json:"type,omitempty"`
	Data XappJWTAuthorizeResponseData `json:"data,omitempty"`
}

type XappJWTAuthorizeResponseData struct {
	Payload string `json:"payload,omitempty"`
}

type XappJWTAuthorizeResponseApp struct {
	Name string `json:"name,omitempty"`
}
