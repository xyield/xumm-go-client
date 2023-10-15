package jwt

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	JWTPINGENDPOINT = "/platform/jwt/ping"
)

// Ping method tests connectivity to XUMM api.
func (j *Jwt) Ping(jwt ...string) (*models.Pong, error) {
	req, err := http.NewRequest(http.MethodGet, j.Cfg.BaseURL+JWTPINGENDPOINT, nil)

	if jwt == nil {
		req.Header = j.Cfg.GetHeaders()
	} else {
		req.Header.Add("Authorization", jwt[0])
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	res, err := j.Cfg.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var p models.Pong
	_, err = utils.DeserialiseRequest(&p, res.Body)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
