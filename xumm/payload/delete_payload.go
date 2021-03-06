package payload

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

// CancelPayloadByUUID cancels a payload, so a user cannot open it anymore.
// Takes a single argument of a payload uuid string.
func (p *Payload) CancelPayloadByUUID(uuid string) (*models.XummDeletePayloadResponse, error) {
	req, err := http.NewRequest(http.MethodDelete, p.Cfg.BaseURL+PAYLOADENDPOINT+uuid, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header = p.Cfg.GetHeaders()

	res, err := p.Cfg.HTTPClient.Do(req)
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
