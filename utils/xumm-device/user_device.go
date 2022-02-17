package xummdevice

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/xyield/xumm-go-client/utils"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	XUMM_API_PREFIX = "https://xumm.app/api/v1/app/"
)

var client http.Client

type UserDevice struct {
	AccessToken            string
	UniqueDeviceIdentifier string
}

type PayloadNotFoundError struct {
	UUID string
}

func (e *PayloadNotFoundError) Error() string {
	return fmt.Sprintf("Payload with UUID %v not found", e.UUID)
}

func NewUserDevice(t, udi string) *UserDevice {
	return &UserDevice{
		AccessToken:            t,
		UniqueDeviceIdentifier: udi,
	}
}

func (u *UserDevice) Ping() (*Pong, error) {
	req, err := http.NewRequest(http.MethodPost, XUMM_API_PREFIX+"ping", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+u.generateBearerToken(strconv.FormatInt(time.Now().UnixNano(), 10)))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var p Pong
	json.Unmarshal(b, &p)

	return &p, nil
}

func (u *UserDevice) OpenPayload(uuid string) error {
	req, err := http.NewRequest(http.MethodGet, XUMM_API_PREFIX+"payload/"+uuid, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+u.generateBearerToken(strconv.FormatInt(time.Now().UnixNano(), 10)))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var p models.XummPayload
	json.Unmarshal(b, &p)

	utils.PrettyPrintJson(p)

	if !p.Meta.Exists {
		return &PayloadNotFoundError{UUID: uuid}
	}

	return nil
}

func (u *UserDevice) SignRequest(uuid string) error {
	ops := []byte(`[
		{"op": "replace","path":"/meta/signed","value": true}
	]`)
	req, err := http.NewRequest(http.MethodPatch, XUMM_API_PREFIX+"payload/"+uuid, bytes.NewBuffer(ops))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+u.generateBearerToken(strconv.FormatInt(time.Now().UnixNano(), 10)))

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var j anyjson.AnyJson
	json.Unmarshal(b, &j)
	utils.PrettyPrintJson(j)
	return nil
}

func (u *UserDevice) generateBearerToken(uid string) string {
	h := sha256.Sum256([]byte(u.AccessToken + u.UniqueDeviceIdentifier + uid))
	s := hex.EncodeToString(h[:])
	return strings.Join([]string{u.AccessToken, uid, s}, ".")
}
