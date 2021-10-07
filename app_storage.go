package xumm

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/models"
)

const (
	APPSTORAGEENDPOINT = "/platform/app-storage"
)

func (c *SDK) GetAppStorage() (*models.AppStorageResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+APPSTORAGEENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	c.SetXummHeaders(req)

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = checkForErrorResponse(res)
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

func (c *SDK) SetAppStorage(d map[string]interface{}) (*models.AppStorageResponse, error) {
	reqBody, err := jsoniter.Marshal(d)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, c.BaseURL+APPSTORAGEENDPOINT, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	c.SetXummHeaders(req)
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	err = checkForErrorResponse(res)
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

func (c *SDK) DeleteAppStorage() (*models.AppStorageResponse, error) {
	req, err := http.NewRequest(http.MethodDelete, c.BaseURL+APPSTORAGEENDPOINT, nil)
	if err != nil {
		return nil, err
	}
	c.SetXummHeaders(req)

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = checkForErrorResponse(res)
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
