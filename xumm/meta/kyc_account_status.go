package meta

import (
	"bytes"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/models"
	"github.com/xyield/xumm-go-client/pkg/utils"
	"github.com/xyield/xumm-go-client/xumm"
)

const (
	KYCACCOUNTSTATUSENDPOINT = "/platform/kyc-status/"
)

// Get account status by xrp public address
func (m *Meta) KycAccountStatus(a string) (*models.KycAccountStatusResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+KYCACCOUNTSTATUSENDPOINT+a, nil)
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

	var kyc models.KycAccountStatusResponse
	_, err = utils.DeserialiseRequest(&kyc, res.Body)
	if err != nil {
		return nil, err
	}

	return &kyc, nil
}

//Get account status by user token body
func (m *Meta) KycStatusState(body models.KycStatusStateRequest) (*models.KycStatusStateResponse, error) {
	reqBody, err := jsoniter.Marshal(body)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, m.Cfg.BaseURL+KYCACCOUNTSTATUSENDPOINT, bytes.NewReader(reqBody))
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

	var kyc models.KycStatusStateResponse
	_, err = utils.DeserialiseRequest(&kyc, res.Body)
	if err != nil {
		return nil, err
	}

	return &kyc, nil
}
