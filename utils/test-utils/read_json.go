package testutils

import (
	"io/ioutil"
	"log"
	"os"
)

func ConvertJsonFileToJsonString(fn string) string {
	// #nosec G304 -- Ignore as this is only used for tests
	f, err := os.Open(fn)

	if err != nil {
		log.Panicln(err)
	}

	// #nosec G307 -- Ignore as this is only used for tests
	defer f.Close()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		log.Panicln(err)
	}

	return string(b)
}
