package meta

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	HOOKHASHENDPOINT = "/platform/hookhash/"
)

type EmptyHookHash struct{}

func (e *EmptyHookHash) Error() string {
	return fmt.Sprintln("No hookhash provided")
}

type InvalidHookHash struct{}

func (e *InvalidHookHash) Error() string {
	return fmt.Sprintln("Invalid hookhash provided, must be 64 hexadecimal characters")
}

func (m *Meta) GetHookHash(h string) (*models.HookHashResponse, error) {

	if len(h) != 64 {
		return nil, &InvalidHookHash{}
	}

	_, err := hex.DecodeString(h)
	if err != nil {
		return nil, &InvalidHookHash{}
	}

	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+HOOKHASHENDPOINT+h, nil)
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
	var hh models.HookHashResponse
	_, err = utils.DeserialiseRequest(&hh, res.Body)
	if err != nil {
		return nil, err
	}

	return &hh, nil
}
