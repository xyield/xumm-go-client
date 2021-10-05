package xumm

import (
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/models"
)

const (
	RATESCURRENCYENDPOINT = "/platform/rates/"
)

func (c *SDK) RatesCurrency(cur string) (*models.RatesCurrencyResponse, error) {

	// check only 3 char cur - regex
	// return error if it is not exactly 3

	// check it is all capitalized

	req, err := http.NewRequest(http.MethodGet, c.BaseURL+RATESCURRENCYENDPOINT+cur, nil)

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

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rc models.RatesCurrencyResponse

	if err = jsoniter.Unmarshal(b, &rc); err != nil {
		log.Println(err)
		return nil, err
	}

	return &rc, nil
}
