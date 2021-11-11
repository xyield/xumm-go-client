package xumm

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
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

	// Possible combos: {ref, code} {ref, message} {ref, code, message}

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
		var e ErrorNotFound

		DeserialiseRequest(&e, res.Body)
		return &e
	}

	var e ErrorResponse

	DeserialiseRequest(&e, res.Body)
	return &e
}

func DeserialiseRequest(v interface{}, body io.ReadCloser) interface{} {
	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
	}

	err = jsoniter.Unmarshal(b, &v)
	if err != nil {
		log.Println(err)
	}
	return v
}
