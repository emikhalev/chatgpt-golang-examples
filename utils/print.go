package utils

import (
	"encoding/json"
)

func SPrintStruct(s interface{}) string {
	json, _ := json.MarshalIndent(s, "", " ")
	return string(json)
}
