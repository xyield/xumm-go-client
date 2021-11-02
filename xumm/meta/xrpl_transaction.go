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
	XRPLTRANSACTIONENDPOINT = "/platform/xrpl-tx/"
)

func (m *Meta) XrplTransaction(txid string) (*models.XrpTxResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+XRPLTRANSACTIONENDPOINT+txid, nil)
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
