package meta

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
	GetKycStatusByAccountENDPOINT = "/platform/kyc-status/"
)

// GetKycStatusByAccount fetches the KYC status for a XUMM user (based on a public XRPL account address, r...).
// Takes 1 parameter, account.
func (m *Meta) GetKycStatusByAccount(a string) (*models.GetKycStatusByAccountResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+GetKycStatusByAccountENDPOINT+a, nil)
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

	var kyc models.GetKycStatusByAccountResponse
	_, err = utils.DeserialiseRequest(&kyc, res.Body)
	if err != nil {
		return nil, err
	}

	return &kyc, nil
}

// GetKycStatusByUserToken fetches the KYC status for a XUMM user (based on an issued user_token).
// Takes 1 parameter, user_token.
func (m *Meta) GetKycStatusByUserToken(body models.GetKycStatusByUserTokenRequest) (*models.GetKycStatusByUserTokenResponse, error) {
	reqBody, err := jsoniter.Marshal(body)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, m.Cfg.BaseURL+GetKycStatusByAccountENDPOINT, bytes.NewReader(reqBody))
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

	var kyc models.GetKycStatusByUserTokenResponse
	_, err = utils.DeserialiseRequest(&kyc, res.Body)
	if err != nil {
		return nil, err
	}

	return &kyc, nil
}
