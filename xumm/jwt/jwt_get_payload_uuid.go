package jwt

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
	"github.com/xyield/xumm-go-client/xumm/payload"
)

// GetPayloadByUUID returns the payload details or payload resolve status and result data.
// Takes a single argument of a payload uuid string.
func (j *Jwt) GetPayloadByUUID(uuid string, jwt ...string) (*models.XummPayload, error) {

	if uuid == "" {
		return nil, &payload.EmptyUuidError{}
	}

	return GetPayload(j, JWTPAYLOADENDPOINT+"/"+uuid, jwt[0])
}

func GetPayload(j *Jwt, endpt string, jwt string) (*models.XummPayload, error) {
	req, err := http.NewRequest(http.MethodGet, j.Cfg.BaseURL+endpt, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if jwt == "" {
		req.Header = j.Cfg.GetHeaders()
	} else {
		req.Header.Add("Authorization", jwt)
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

	var pur models.XummPayload

	_, err = utils.DeserialiseRequest(&pur, res.Body)
	if err != nil {
		return nil, err
	}

	return &pur, nil
}
