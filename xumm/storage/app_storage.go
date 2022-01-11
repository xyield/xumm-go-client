package storage

import (
	"bytes"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client"
	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm/models"
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
	req.Header = s.Cfg.GetHeaders()

	res, err := s.Cfg.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}

	var as models.AppStorageResponse
	_, err = utils.DeserialiseRequest(&as, res.Body)
	if err != nil {
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
	req.Header = s.Cfg.GetHeaders()
	res, err := s.Cfg.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}

	var as models.AppStorageResponse
	_, err = utils.DeserialiseRequest(&as, res.Body)
	if err != nil {
		return nil, err
	}

	return &as, nil
}

func (s *Storage) DeleteAppStorage() (*models.AppStorageResponse, error) {
	req, err := http.NewRequest(http.MethodDelete, s.Cfg.BaseURL+APPSTORAGEENDPOINT, nil)
	if err != nil {
		return nil, err
	}

	req.Header = s.Cfg.GetHeaders()
	res, err := s.Cfg.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}

	var as models.AppStorageResponse
	_, err = utils.DeserialiseRequest(&as, res.Body)
	if err != nil {
		return nil, err
	}

	return &as, nil
}
