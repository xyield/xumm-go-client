package meta

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	XRPLTRANSACTIONENDPOINT = "/platform/xrpl-tx/"
)

// GetXrplTransaction fetches transaction & outcome live from XRP ledger full history nodes (through the XUMM platform) containing parsed transaction outcome balance mutations.
// Takes 1 parameter, txid (64 hexadecimal characters).
func (m *Meta) GetXrplTransaction(txid string) (*models.XrpTxResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+XRPLTRANSACTIONENDPOINT+txid, nil)
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

	var tx models.XrpTxResponse
	_, err = utils.DeserialiseRequest(&tx, res.Body)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}
