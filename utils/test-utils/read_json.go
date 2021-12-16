package testutils

import (
	"io/ioutil"
	"log"
	"os"
)

func ConvertJsonFileToJsonString(fn string) string {
	f, err := os.Open(fn)

	if err != nil {
		log.Panicln(err)
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		log.Panicln(err)
	}

	return string(b)
}
