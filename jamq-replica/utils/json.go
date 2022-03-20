package utils

import (
	"encoding/json"
	"jamq-replica/globals"
)

func GetRequestBody(requestBody interface{}) error {
	err := json.Unmarshal([]byte(globals.Body), requestBody)
	return err
}
