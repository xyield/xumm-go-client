package json

import (
	"github.com/ugorji/go/codec"
)

type AnyJson map[string]interface{}

func (a *AnyJson) UnmarshalJSON(data []byte) error {

	var jh codec.JsonHandle

	jh.SignedInteger = true

	err := codec.NewDecoderBytes(data, &jh).Decode(a)
	return err
}
