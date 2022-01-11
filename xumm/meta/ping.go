package meta

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client"
	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	PINGENDPOINT = "/platform/ping"
)

func (m *Meta) Ping() (*models.Pong, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+PINGENDPOINT, nil)
	req.Header = m.Cfg.GetHeaders()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := m.Cfg.HTTPClient.Do(req)
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
