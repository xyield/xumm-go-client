package json

import (
	"reflect"

	"github.com/ugorji/go/codec"
)

type AnyJson map[string]interface{}

func (a *AnyJson) UnmarshalJSON(data []byte) error {

	var jh codec.JsonHandle

	jh.SignedInteger = true
	jh.MapType = reflect.TypeOf(map[string]interface{}{})

	err := codec.NewDecoderBytes(data, &jh).Decode(a)

	return err
}
