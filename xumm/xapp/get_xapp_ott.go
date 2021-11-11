package xapp

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/pkg/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

// type invalidToken struct{}

// func (e *invalidToken) Error() string {
// 	return fmt.Sprintln("Invalid OTT entered")
// }

const (
	GETXAPPOTTENDPOINT = "/platform/xapp/ott/"
)

func (x *Xapp) GetXappOtt(t string) (*models.XappResponse, error) {

	req, err := http.NewRequest(http.MethodGet, x.Cfg.BaseURL+GETXAPPOTTENDPOINT+t, nil)
	req.Header = x.Cfg.Headers
	if err != nil {
		log.Println(err)
		return nil, err
	}
	res, err := x.Cfg.HTTPClient.Do(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var xr models.XappResponse
	_, err = utils.DeserialiseRequest(&xr, res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &xr, err
}
