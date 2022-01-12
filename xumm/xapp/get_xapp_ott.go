package xapp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

// InvalidToken returns an error when an invalid one time token is given.
type InvalidToken struct{}

func (e *InvalidToken) Error() string {
	return fmt.Sprintln("Invalid OTT entered")
}

const (
	XAPPENDPOINT = "/platform/xapp/"
)
// GetXappOtt allows the xApp to retrieve verified session related information from the XUMM user.
// xApps are embedded apps. Publishing an xApp and calling xApp API endpoints are only available for XRPL Labs / XUMM partners.
// 1 parameter is required, a token (UUID).
func (x *Xapp) GetXappOtt(t string) (*models.XappOttResponse, error) {

	if t == "" {
		return nil, &InvalidToken{}
	}

	req, err := http.NewRequest(http.MethodGet, x.Cfg.BaseURL+XAPPENDPOINT+"ott/"+t, nil)
	req.Header = x.Cfg.GetHeaders()
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

	var xr models.XappOttResponse
	_, err = utils.DeserialiseRequest(&xr, res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &xr, err
}
