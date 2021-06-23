package xumm

import (
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/models"
)

const (
	XRPLTRANSACTIONENDPOINT = "platform/xrpl-tx/"
)

func (c *SDK) XrplTransaction(txid string) (*models.XrpTxResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+XRPLTRANSACTIONENDPOINT+txid, nil)
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

	var tx models.XrpTxResponse

	if err = jsoniter.Unmarshal(b, &tx); err != nil {
		log.Println(err)
		return nil, err
	}

	return &tx, nil
}
