package xumm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gitlab.com/xyield/xumm-go-client/models"
)

const (
	PINGENDPOINT = "/platform/ping"
)

func (c *Client) Ping() *models.Pong {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+PINGENDPOINT, nil)
	if err != nil {
		log.Println(err)
		return &models.Pong{}
	}
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		log.Println(err)
	}
	var p models.Pong

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return &models.Pong{}
	}

	if err = json.Unmarshal(b, &p); err != nil {
		log.Println(err)
		return &models.Pong{}
	}

	return &p
}
