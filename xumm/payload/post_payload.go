package payload

import (
	"bytes"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/pkg/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	POSTPAYLOADENDPOINT = "/platform/payload"
)

type TransactionTypeError struct {
	Code string
}

func (e *TransactionTypeError) Error() string {
	return "Invalid transaction type provided."
}

func (p *Payload) PostPayload(body models.XummPostPayload) (*models.XummPostPayloadResponse, error) {

	reqBody, err := jsoniter.Marshal(body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if body.TxJson != nil {
		tx, err := jsoniter.Marshal(body.TxJson)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		var jt models.JsonTransaction
		_, err = utils.DeserialiseRequest(&jt, bytes.NewReader(tx))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if jt.TransactionType == "" {
			log.Println("No valid transaction type provided in TxJson")
			return nil, &TransactionTypeError{}
		}
	}

	req, err := http.NewRequest(http.MethodGet, p.Cfg.BaseURL+POSTPAYLOADENDPOINT, bytes.NewReader(reqBody))

	req.Header = p.Cfg.Headers
	if err != nil {
		log.Println(err)
		return nil, err
	}
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

	var pr models.XummPostPayloadResponse
	_, err = utils.DeserialiseRequest(&pr, res.Body)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}
