package v1

import "github.com/labstack/echo/v4"

func HandleV1Routes(group *echo.Group) {
	group.POST("/queues", HandleCreateQueue)
}
