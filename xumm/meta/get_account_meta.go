package meta

import (
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	ACCOUNTMETAENDPOINT = "/platform/account-meta/"
)

func (m *Meta) GetAccountMeta(a string) (*models.AccountMetaResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+ACCOUNTMETAENDPOINT+a, nil)
	if err != nil {
		return nil, err
	}
	req.Header = m.Cfg.GetHeaders()
	res, err := m.Cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var am models.AccountMetaResponse

	_, err = utils.DeserialiseRequest(&am, res.Body)
	if err != nil {
		return nil, err
	}

	return &am, nil
}
