package xumm

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xyield/xumm-go-client/pkg/utils"
)

type ErrorResponse struct {
	ErrorResponseInternal ErrorResponseInternal `json:"error"`
}

type ErrorResponseInternal struct {
	Reference string `json:"reference"`
	Code      int    `json:"code"`
}

type ErrorUnauthorised struct {
	ErrorTest bool   `json:"error"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
	Code      int    `json:"code"`
	Req       string `json:"req"`
	Method    string `json:"method"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Error returned with reference %v and code %v", e.ErrorResponseInternal.Reference, e.ErrorResponseInternal.Code)
}

func (e *ErrorUnauthorised) Error() string {
	return fmt.Sprintf("Error returned with code %v and message '%v'", e.Code, e.Message)
}

func CheckForErrorResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	log.Println("Error response recieved from Xumm")

	if res.StatusCode == 404 {
		var e ErrorUnauthorised

		utils.DeserialiseRequest(&e, res.Body)
		return &e
	}

	var e ErrorResponse

	_, err := utils.DeserialiseRequest(&e, res.Body)
	if err != nil {
		return err
	}

	return &e
}
