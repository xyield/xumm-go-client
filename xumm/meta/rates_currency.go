package meta

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/models"
	"github.com/xyield/xumm-go-client/xumm"
)

const (
	RATESCURRENCYENDPOINT = "/platform/rates/"
)

type CurrencyCodeError struct {
	Code string
}

func (e *CurrencyCodeError) Error() string {
	return fmt.Sprintf("Currency code %v is not valid", e.Code)
}

func (m *Meta) RatesCurrency(cur string) (*models.RatesCurrencyResponse, error) {

	ok, _ := regexp.MatchString(`^[a-zA-Z]{3}$`, cur)

	if !ok {
		return nil, &CurrencyCodeError{Code: cur}
	}

	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+RATESCURRENCYENDPOINT+cur, nil)

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

	var rc models.RatesCurrencyResponse

	if err = jsoniter.Unmarshal(b, &rc); err != nil {
		log.Println(err)
		return nil, err
	}

	return &rc, nil
}
