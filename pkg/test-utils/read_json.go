package testutils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	str := string(b)
	str = strings.Replace(str, "\n", "", -1)
	// TODO: find a safer way to parse out spaces
	str = strings.Replace(str, " ", "", -1)

	return str
}
