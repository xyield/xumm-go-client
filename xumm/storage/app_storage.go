package storage

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/models"
	"github.com/xyield/xumm-go-client/xumm"
)

const (
	APPSTORAGEENDPOINT = "/platform/app-storage"
)

type StorageInterface interface {
	GetAppStorage() (*models.AppStorageResponse, error)
	SetAppStorage(d map[string]interface{}) (*models.AppStorageResponse, error)
	DeleteAppStorage() (*models.AppStorageResponse, error)
}

type Storage struct {
	Cfg *xumm.Config
}

func (s *Storage) GetAppStorage() (*models.AppStorageResponse, error) {
	req, err := http.NewRequest(http.MethodGet, s.Cfg.BaseURL+APPSTORAGEENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	req.Header = map[string][]string{
		"XUMM_API_KEY":    {s.Cfg.ApiKey},
		"XUMM_API_SECRET": {s.Cfg.ApiSecret},
		"Content-Type":    {"application/json"},
	}

	res, err := s.Cfg.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}
	var as models.AppStorageResponse

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err = jsoniter.Unmarshal(b, &as); err != nil {
		return nil, err
	}

	return &as, nil
}

func (s *Storage) SetAppStorage(d map[string]interface{}) (*models.AppStorageResponse, error) {
	reqBody, err := jsoniter.Marshal(d)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, s.Cfg.BaseURL+APPSTORAGEENDPOINT, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header = map[string][]string{
		"XUMM_API_KEY":    {s.Cfg.ApiKey},
		"XUMM_API_SECRET": {s.Cfg.ApiSecret},
		"Content-Type":    {"application/json"},
	}
	res, err := s.Cfg.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var as models.AppStorageResponse

	if err = jsoniter.Unmarshal(resBody, &as); err != nil {
		return nil, err
	}

	return &as, nil
}

func (s *Storage) DeleteAppStorage() (*models.AppStorageResponse, error) {
	req, err := http.NewRequest(http.MethodDelete, s.Cfg.BaseURL+APPSTORAGEENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	req.Header = map[string][]string{
		"XUMM_API_KEY":    {s.Cfg.ApiKey},
		"XUMM_API_SECRET": {s.Cfg.ApiSecret},
		"Content-Type":    {"application/json"},
	}
	res, err := s.Cfg.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}
	var as models.AppStorageResponse

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err = jsoniter.Unmarshal(b, &as); err != nil {
		return nil, err
	}

	return &as, nil
}
