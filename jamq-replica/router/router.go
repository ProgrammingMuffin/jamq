package router

import (
	"jamq-replica/router/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Authenticate)
	return e
}
