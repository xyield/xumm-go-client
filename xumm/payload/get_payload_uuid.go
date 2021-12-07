package payload

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/pkg/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

type EmptyUuidError struct {
}

func (e *EmptyUuidError) Error() string {
	return "Empty UUID provided."
}

func (p *Payload) GetPayloadByUUID(uuid string) (*models.XummPayload, error) {

	if uuid == "" {
		return nil, &EmptyUuidError{}
	}

	return GetPayload(p, PAYLOADENDPOINT+uuid)
}

func GetPayload(p *Payload, endpt string) (*models.XummPayload, error) {
	req, err := http.NewRequest(http.MethodGet, p.Cfg.BaseURL+endpt, nil)
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

	var pur models.XummPayload

	_, err = utils.DeserialiseRequest(&pur, res.Body)
	if err != nil {
		return nil, err
	}

	return &pur, nil
}
