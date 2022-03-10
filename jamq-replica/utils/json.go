package utils

import (
	"bytes"
	"encoding/json"

	"github.com/labstack/echo/v4"
)

func GetRequestBody(c echo.Context, requestBody interface{}) error {
	bytesBuff := new(bytes.Buffer)
	bytesBuff.ReadFrom(c.Request().Body)
	err := json.Unmarshal(bytesBuff.Bytes(), requestBody)
	return err
}
