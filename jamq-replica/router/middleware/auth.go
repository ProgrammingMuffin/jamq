package middleware

import (
	"encoding/base64"
	"fmt"
	"jamq-replica/constants"
	"jamq-replica/globals"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"crypto/hmac"
	"crypto/sha256"

	"github.com/labstack/echo/v4"
	"github.com/valyala/bytebufferpool"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bodyBytesBuffer := &bytebufferpool.ByteBuffer{}
		request := c.Request()
		header := request.Header
		bodyBytesBuffer.ReadFrom(request.Body)
		body := bodyBytesBuffer.String()
		globals.Body = body
		fmt.Println("Body is: " + body)
		queryString := strings.ToValidUTF8(c.QueryString(), "")
		nonce := header.Get(constants.GetHeaders().X_AUTH_NONCE)
		sentTimestamp, err := strconv.Atoi(nonce)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid nonce, error: "+err.Error())
		}
		signature := header.Get(constants.GetHeaders().X_AUTH_SIGNATURE)
		currentTimestamp := time.Now().UnixMilli()
		configNonceTTL, err := strconv.Atoi(os.Getenv(constants.GetConfig().NONCE_TTL))
		if (currentTimestamp - int64(sentTimestamp)) > int64(configNonceTTL) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Nonce expired")
		}
		expectedSignature := createSignature(request.Method, request.URL.Path, body,
			queryString, nonce, os.Getenv(constants.GetConfig().SECRET))
		fmt.Println("signature is: " + expectedSignature)
		if signature != expectedSignature {
			return echo.NewHTTPError(http.StatusUnauthorized, "Signature doesn't match")
		}
		err = next(c)
		return err
	}
}

func createSignature(method string, uri string, body string, queryString string, nonce string, signingKey string) string {
	signingString := "uri=" + uri + "&method=" + method + "&nonce=" + nonce + queryString
	if method == http.MethodPost || method == http.MethodPut {
		signingString += body
	}
	fmt.Println("signing string is: " + signingString)
	mac := hmac.New(sha256.New, []byte(signingKey))
	mac.Write([]byte(signingString))
	return strings.ToValidUTF8(base64.StdEncoding.EncodeToString(mac.Sum([]byte(nil))), "")
}
