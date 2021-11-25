package utils

import (
	"io"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
)

func DeserialiseRequest(v interface{}, body io.Reader) (interface{}, error) {

	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return UnmarshalResponse(v, b)
}

func UnmarshalResponse(v interface{}, b []byte) (interface{}, error) {

	err := jsoniter.Unmarshal(b, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return v, nil
}
