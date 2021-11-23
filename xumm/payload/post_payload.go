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

func (p *Payload) PostPayload(body models.XummPostPayload) (*models.XummPostPayloadResponse, error) {

	if body.TxJson != nil {
		if _, ok := body.TxJson["TransactionType"]; !ok {
			return nil, &TransactionTypeError{}
		}
		tx := body.TxJson["TransactionType"]
		if _, ok := transactionTypeFromString[tx.(string)]; !ok {
			return nil, &TransactionTypeError{}
		}
		// tx, err := jsoniter.Marshal(body.TxJson)
		// if err != nil {
		// 	log.Println(err)
		// 	return nil, err
		// }
		// var jt models.JsonTransaction
		// _, err = utils.DeserialiseRequest(&jt, bytes.NewReader(tx))
		// if err != nil {
		// 	log.Println(err)
		// 	return nil, err
		// }
		// if jt.TransactionType == "" {
		// 	log.Println("No valid transaction type provided in TxJson")
		// 	return nil, &TransactionTypeError{}
		// }
	}

	reqBody, err := jsoniter.Marshal(body)
	if err != nil {
		log.Println(err)
		return nil, err
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
