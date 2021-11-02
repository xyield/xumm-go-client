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
	CURATEDASSETSENDPOINT = "/platform/curated-assets"
)

func (m *Meta) CuratedAssets() (*models.CuratedAssetsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+CURATEDASSETSENDPOINT, nil)
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
	var ca models.CuratedAssetsResponse

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = jsoniter.Unmarshal(b, &ca); err != nil {
		log.Println(err)
		return nil, err
	}

	return &ca, nil
}
