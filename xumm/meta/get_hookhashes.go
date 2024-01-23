package meta

import (
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

func (m *Meta) GetHookHashes() (*models.HookHashesResponse, error) {

	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+HOOKHASHENDPOINT, nil)
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
	var hh models.HookHashesResponse
	_, err = utils.DeserialiseRequest(&hh, res.Body)
	if err != nil {
		return nil, err
	}

	return &hh, nil
}
