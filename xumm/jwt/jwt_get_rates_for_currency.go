package jwt

import (
	"log"
	"net/http"
	"regexp"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/meta"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	JWTRATESCURRENCYENDPOINT = "/jwt/rates/"
)

// GetRatesForCurrency gets semi-live XRP exchange rates.
// Takes 1 parameter, a 3 alpha char currency code, eg. INR.
func (j *Jwt) GetRatesForCurrency(cur string, jwt ...string) (*models.RatesCurrencyResponse, error) {

	ok, _ := regexp.MatchString(`^[a-zA-Z]{3}$`, cur)

	if !ok {
		return nil, &meta.CurrencyCodeError{Code: cur}
	}

	req, err := http.NewRequest(http.MethodGet, j.Cfg.BaseURL+JWTRATESCURRENCYENDPOINT+cur, nil)

	if jwt == nil {
		req.Header = j.Cfg.GetHeaders()
	} else {
		req.Header.Add("Authorization", jwt[0])
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	res, err := j.Cfg.HTTPClient.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var rc models.RatesCurrencyResponse
	_, err = utils.DeserialiseRequest(&rc, res.Body)
	if err != nil {
		return nil, err
	}

	return &rc, nil
}
