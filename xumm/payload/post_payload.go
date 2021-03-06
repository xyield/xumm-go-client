package payload

import (
	"bytes"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	PAYLOADENDPOINT = "/platform/payload"
)

// PostPayload method submits a payload containing a sign request to the XUMM platform.
// Takes a single argument of a XummPostPayload which MUST have either TxJson or TxBlob field filled.
// If using TxJson, it must be formatted as per XRPL transaction format specification at xrpl.org.
func (p *Payload) PostPayload(body models.XummPostPayload) (*models.CreatedPayload, error) {

	if body.TxJson != nil {
		if _, ok := body.TxJson["TransactionType"]; !ok {
			return nil, &TransactionTypeError{}
		}
		tx := body.TxJson["TransactionType"]
		if _, ok := transactionTypeFromString[tx.(string)]; !ok {
			return nil, &TransactionTypeError{}
		}
	}

	reqBody, err := jsoniter.Marshal(body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, p.Cfg.BaseURL+PAYLOADENDPOINT, bytes.NewReader(reqBody))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header = p.Cfg.GetHeaders()

	res, err := p.Cfg.HTTPClient.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var pr models.CreatedPayload
	_, err = utils.DeserialiseRequest(&pr, res.Body)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}
