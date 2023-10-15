package meta

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	CURATEDASSETSENDPOINT = "/platform/curated-assets"
)

// GetCuratedAssets Gets curated assets from the XUMM API. This API contains the same issuers and assets available to users in XUMM when they press the "Add asset" button on the home screen.
func (m *Meta) GetCuratedAssets() (*models.CuratedAssetsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+CURATEDASSETSENDPOINT, nil)
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
	var ca models.CuratedAssetsResponse

	_, err = utils.DeserialiseRequest(&ca, res.Body)
	if err != nil {
		return nil, err
	}

	return &ca, nil
}
