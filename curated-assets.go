package xumm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gitlab.com/xYield/xumm-go-client/models"
)

const (
	CURATEDASSETSENDPOINT = "/curated-assets"
)

func (c *SDK) CurratedAssets() (*models.CurratedAssetsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+CURATEDASSETSENDPOINT, nil)
	c.SetXummHeaders(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = checkForErrorResponse(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var ca models.CurratedAssetsResponse

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = json.Unmarshal(b, &ca); err != nil {
		log.Println(err)
		return nil, err
	}

	return &ca, nil
}
