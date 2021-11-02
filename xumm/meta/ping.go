package meta

import (
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/models"
	"github.com/xyield/xumm-go-client/xumm"
)

const (
	PINGENDPOINT = "/platform/ping"
)

func (m *Meta) Ping() (*models.Pong, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+PINGENDPOINT, nil)
	req.Header = m.Cfg.Headers
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

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = jsoniter.Unmarshal(b, &p); err != nil {
		log.Println(err)
		return nil, err
	}

	return &p, nil
}
