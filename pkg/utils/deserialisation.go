package utils

import (
	"io"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
)

func DeserialiseRequest(v interface{}, body io.ReadCloser) (interface{}, error) {

	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = jsoniter.Unmarshal(b, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return v, nil
}
