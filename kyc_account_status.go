package xumm

import (
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/models"
)

const (
	KYCACCOUNTSTATUSENDPOINT = "/platform/kyc-status/"
)

func (c *SDK) KycAccountStatus(a string) (*models.KycAccountStatusResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+KYCACCOUNTSTATUSENDPOINT+a, nil)
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

	var kyc models.KycAccountStatusResponse

	if err = jsoniter.Unmarshal(b, &kyc); err != nil {
		log.Println(err)
		return nil, err
	}

	return &kyc, nil
}
