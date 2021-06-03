package xumm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gitlab.com/xYield/xumm-go-client/models"
)

const (
	PINGENDPOINT = "/platform/ping"
)

func (c *SDK) Ping() (*models.Pong, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+PINGENDPOINT, nil)
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
	var p models.Pong

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = json.Unmarshal(b, &p); err != nil {
		log.Println(err)
		return nil, err
	}

	return &p, nil
}
