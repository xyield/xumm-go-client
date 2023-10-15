package jwt

import (
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	JWTAUTHORIZEENDPOINT = "/jwt/authorize"
)

func (j *Jwt) JWTGetAuthorization() (*models.XappJWTAuthorizeResponse, error) {
	req, err := http.NewRequest(http.MethodGet, j.Cfg.BaseURL+JWTAUTHORIZEENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	req.Header = j.Cfg.GetHeaders()
	res, err := j.Cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}
	var jr models.XappJWTAuthorizeResponse
	_, err = utils.DeserialiseRequest(&jr, res.Body)
	if err != nil {
		return nil, err
	}

	return &jr, nil
}
