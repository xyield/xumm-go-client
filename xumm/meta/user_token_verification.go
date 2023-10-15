package meta

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/xyield/xumm-go-client/utils"
	"github.com/xyield/xumm-go-client/xumm"
	"github.com/xyield/xumm-go-client/xumm/models"
)

const (
	USERTOKENENDPOINT = "/platform/user-token"
)

// EmptyUserToken Returns an error if no user tokens are provided.
type EmptyUserToken struct {
	multiple bool
}

func (e *EmptyUserToken) Error() string {
	if e.multiple {
		return fmt.Sprintln("No user tokens provided")
	}
	return fmt.Sprintln("No user token provided")
}

// Verifys a single user token.
// Takes a single argument of a user token.
func (m *Meta) VerifyUserToken(t string) (*models.UserTokenResponse, error) {

	if t == "" {
		return nil, &EmptyUserToken{}
	}

	path := strings.Join([]string{USERTOKENENDPOINT, t}, "/")
	req, err := http.NewRequest(http.MethodGet, m.Cfg.BaseURL+path, nil)
	req.Header = m.Cfg.GetHeaders()
	if err != nil {
		return nil, err
	}

	res, err := m.Cfg.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = xumm.CheckForErrorResponse(res)
	if err != nil {
		return nil, err
	}

	var uts models.UserTokenResponse
	_, err = utils.DeserialiseRequest(&uts, res.Body)
	if err != nil {
		return nil, err
	}

	return &uts, nil
}

// Verifys multiple user tokens.
// Takes multiple user token strings as arguments.
func (m *Meta) VerifyUserTokens(uts ...string) (*models.UserTokenResponse, error) {

	if len(uts) == 0 {
		return nil, &EmptyUserToken{multiple: true}
	}

	ts := &models.UserTokenRequest{
		Tokens: uts,
	}

	reqBody, err := jsoniter.Marshal(ts)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, m.Cfg.BaseURL+USERTOKENENDPOINT, bytes.NewReader(reqBody))
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

	var userTokens models.UserTokenResponse
	_, err = utils.DeserialiseRequest(&userTokens, res.Body)
	if err != nil {
		return nil, err
	}
	return &userTokens, nil
}
