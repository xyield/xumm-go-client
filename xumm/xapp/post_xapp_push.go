package xapp

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/xyield/xumm-go-client"
	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type invalidPushRequestError struct{}

func (*invalidPushRequestError) Error() string {
	return "Empty user token and/or subtitle provided."
}

func (x *Xapp) PostXappPush(b models.XappRequest) (*models.XappResponse, error) {
	if b.UserToken == "" || b.Subtitle == "" {
		return nil, &invalidPushRequestError{}
	}

	reqBody, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, x.Cfg.BaseURL+XAPPENDPOINT+"push", bytes.NewReader(reqBody))
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

	var push models.XappResponse

	_, err = utils.DeserialiseRequest(&push, res.Body)
	if err != nil {
		return nil, err
	}

	return &push, nil

}
