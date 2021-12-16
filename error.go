package xumm

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/utils"
	anyjson "github.com/xyield/xumm-go-client/utils/json"
)

type ErrorResponse struct {
	ErrorResponseBody ErrorResponseBody `json:"error"`
}

type ErrorResponseBody struct {
	Reference string `json:"reference"`
	Code      int    `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
}

type ErrorNotFound struct {
	Err       bool   `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
	Reference string `json:"reference"`
	Code      int    `json:"code"`
	Req       string `json:"req,omitempty"`
	Method    string `json:"method,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e.ErrorResponseBody.Message != "" && e.ErrorResponseBody.Code != 0 {
		return fmt.Sprintf("Error returned with reference %v, code %v and message '%v'", e.ErrorResponseBody.Reference, e.ErrorResponseBody.Code, e.ErrorResponseBody.Message)
	}

	if e.ErrorResponseBody.Message != "" {
		return fmt.Sprintf("Error returned with reference %v and message '%v'", e.ErrorResponseBody.Reference, e.ErrorResponseBody.Message)
	}

	return fmt.Sprintf("Error returned with reference %v and code %v", e.ErrorResponseBody.Reference, e.ErrorResponseBody.Code)
}

func (e *ErrorNotFound) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("Error returned with code %v, reference '%v' and message '%v'", e.Code, e.Reference, e.Message)
	}
	return fmt.Sprintf("Error returned with code %v and reference '%v'", e.Code, e.Reference)
}

func CheckForErrorResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	log.Println("Error response received from Xumm")

	if res.StatusCode == 404 {

		b, _ := ioutil.ReadAll(res.Body)
		var aj anyjson.AnyJson
		_, err := utils.UnmarshalResponse(&aj, b)
		if err != nil {
			return err
		}

		if _, ok := aj["code"]; ok {

			var e ErrorNotFound

			_, err := utils.UnmarshalResponse(&e, b)
			if err != nil {
				return err
			}
			return &e
		} else {
			var e ErrorResponse

			_, err := utils.UnmarshalResponse(&e, b)
			if err != nil {
				return err
			}
			return &e
		}
	}

	var e ErrorResponse

	_, err := utils.DeserialiseRequest(&e, res.Body)
	if err != nil {
		return err
	}

	return &e
}
