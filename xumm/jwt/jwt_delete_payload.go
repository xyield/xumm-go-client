package jwt

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

// CancelPayloadByUUID cancels a payload, so a user cannot open it anymore.
// Takes a single argument of a payload uuid string.
func (j *Jwt) CancelPayloadByUUID(uuid string, jwt ...string) (*models.XummDeletePayloadResponse, error) {
	req, err := http.NewRequest(http.MethodDelete, j.Cfg.BaseURL+JWTPAYLOADENDPOINT+uuid, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

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

	var dr models.XummDeletePayloadResponse

	_, err = utils.DeserialiseRequest(&dr, res.Body)
	if err != nil {
		return nil, err
	}

	return &dr, nil
}
