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

type PatchResponse struct {
	ReferenceCallUuidv4 string      `json:"reference_call_uuidv4"`
	Signed              bool        `json:"signed"`
	UserToken           interface{} `json:"user_token"`
	ReturnURL           ReturnUrl   `json:"return_url"`
}
type ReturnUrl struct {
	App string `json:"app"`
	Web string `json:"web"`
}

// type RejectPayload struct {
// 	Rejected bool `json:"reject"`
// }
type SignPayload struct {
	// SignedBlob string `json:"signed_blob"`
	// TxID       string `json:"tx_id"`
	// Multisigned string     `json:"multisigned"`
	// Dispatched  Dispatched `json:"dispatched"`
	Permission Permission `json:"permission"`
}
type Dispatched struct {
	To     string `json:"to"`
	Result string `json:"result"`
}
type Permission struct {
	Push bool `json:"push"`
	Days int  `json:"days"`
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
	err = json.Unmarshal(b, &p)
	if err != nil {
		return nil, err
	}

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
	err = json.Unmarshal(b, &p)
	if err != nil {
		return err
	}

	utils.PrettyPrintJson(p)

	if !p.Meta.Exists {
		return &PayloadNotFoundError{UUID: uuid}
	}

	return nil
}

func (u *UserDevice) SignRequest(uuid string) (string, error) {

	// conditional here for multisign with those fields
	s := SignPayload{
		// SignedBlob: "xxxx",
		// TxID:       "yyyy",
		Permission: Permission{
			Push: true,
			Days: 365,
		},
	}

	body, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	req, err := http.NewRequest(http.MethodPatch, XUMM_API_PREFIX+"payload/"+uuid, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+u.generateBearerToken(strconv.FormatInt(time.Now().UnixNano(), 10)))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var j anyjson.AnyJson
	err = json.Unmarshal(b, &j)
	if err != nil {
		return "", err
	}
	utils.PrettyPrintJson(j)

	json, _ := json.MarshalIndent(j, "", "  ")
	return string(json), nil
}

func (u *UserDevice) RejectRequest(uuid string) (string, error) {
	ops := []byte(`{
		"reject": true
	}`)
	req, err := http.NewRequest(http.MethodPatch, XUMM_API_PREFIX+"payload/"+uuid, bytes.NewBuffer(ops))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+u.generateBearerToken(strconv.FormatInt(time.Now().UnixNano(), 10)))

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var j anyjson.AnyJson
	err = json.Unmarshal(b, &j)
	if err != nil {
		return "", err
	}
	utils.PrettyPrintJson(j)

	json, _ := json.MarshalIndent(j, "", "  ")
	return string(json), nil
}

func (u *UserDevice) generateBearerToken(uid string) string {
	h := sha256.Sum256([]byte(u.AccessToken + u.UniqueDeviceIdentifier + uid))
	s := hex.EncodeToString(h[:])
	return strings.Join([]string{u.AccessToken, uid, s}, ".")
}
