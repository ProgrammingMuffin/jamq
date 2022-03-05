package middleware

import (
	"encoding/json"
	"jamq-replica/constants"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"crypto/hmac"
	"crypto/sha256"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return authenticationMiddleware
}

func authenticationMiddleware(c echo.Context) error {
	var bodyBytes []byte
	request := c.Request()
	header := request.Header
	request.Body.Read(bodyBytes)
	body := string(bodyBytes)
	nonce := header.Get(constants.GetHeaders().X_AUTH_NONCE)
	sentTimestamp, err := strconv.Atoi(nonce)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid nonce")
	}
	signed := header.Get(constants.GetHeaders().X_AUTH_SIGNED)
	signature := header.Get(constants.GetHeaders().X_AUTH_SIGNATURE)
	currentTimestamp := time.Now().UnixMilli()
	if (currentTimestamp - int64(sentTimestamp)) > os.Getenv(constants.GetConfig().NONCE_TTL) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Nonce expired")
	}
	expectedSignature := createSignature(request.Method, request.URL.Path, body,
		signed, nonce, os.Getenv(constants.GetConfig().SECRET))
	if signature != expectedSignature {
		return echo.NewHTTPError(http.StatusUnauthorized, "Signature doesn't match")
	}
	return nil
}

func createSignature(method string, uri string, body string, signed string, nonce string, signingKey string) string {
	signingString := "uri=" + uri + "&method=" + method + "&nonce=" + nonce
	if method == http.MethodPost {
		if body != "" {
			var bodyJson map[string]interface{}
			json.Unmarshal([]byte(body), bodyJson)
			for _, key := range strings.Split(body, ",") {
				signingString += "&" + key + "=" + bodyJson[key].(string)
			}
		}
	}
	mac := hmac.New(sha256.New, []byte(signingKey))
	mac.Write([]byte(signingString))
	return string(mac.Sum(nil))
}
