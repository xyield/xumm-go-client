package jwt

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	JWTCURATEDASSETSENDPOINT = "/jwt/curated-assets"
)

// GetCuratedAssets Gets curated assets from the XUMM API. This API contains the same issuers and assets available to users in XUMM when they press the "Add asset" button on the home screen.
func (j *Jwt) GetCuratedAssets(jwt ...string) (*models.CuratedAssetsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, j.Cfg.BaseURL+JWTCURATEDASSETSENDPOINT, nil)
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
	var ca models.CuratedAssetsResponse

	_, err = utils.DeserialiseRequest(&ca, res.Body)
	if err != nil {
		return nil, err
	}

	return &ca, nil
}
