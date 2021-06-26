package xumm

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type ErrorResponse struct {
	ErrorResponseInternal ErrorResponseInternal `json:"error"`
}

type ErrorResponseInternal struct {
	Reference string `json:"reference"`
	Code      int    `json:"code"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Error returned with reference %v and code %v", e.ErrorResponseInternal.Reference, e.ErrorResponseInternal.Code)
}

func checkForErrorResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	log.Println("Error response recieved from Xumm")
	var e ErrorResponse

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	err = jsoniter.Unmarshal(b, &e)
	if err != nil {
		log.Println(err)
	}
	return &e
}
