package xapp

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type invalidEventRequestError struct{}

func (*invalidEventRequestError) Error() string {
	return "Empty user token and/or subtitle provided."
}

// PostXappEvent allows publishing an xApp event in the "Requests" event list of a user, while sending a Push notification pointing to the event.
// 2 parameters are required for the POST request, b.UserToken and b.Subtitle.
func (x *Xapp) PostXappEvent(b models.XappRequest) (*models.XappResponse, error) {

	if b.UserToken == "" || b.Subtitle == "" {
		return nil, &invalidEventRequestError{}
	}

	reqBody, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, x.Cfg.BaseURL+XAPPENDPOINT+"event", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header = x.Cfg.GetHeaders()

	res, err := x.Cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}

	var event models.XappResponse

	_, err = utils.DeserialiseRequest(&event, res.Body)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
