package meta

import (
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	RAILSENDPOINT = "/platform/rails"
)

func (m *Meta) GetRails() (*models.RailsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+RAILSENDPOINT, nil)
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
		return nil, err
	}
	var r models.RailsResponse
	_, err = utils.DeserialiseRequest(&r, res.Body)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
