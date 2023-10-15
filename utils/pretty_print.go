package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrintJson(v interface{}) {
	json, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(json))
}
