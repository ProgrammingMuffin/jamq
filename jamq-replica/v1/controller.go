package v1

import "github.com/labstack/echo/v4"

func HandleV1Routes(group *echo.Group) {
	group.POST("/queues", HandleCreateQueue)
	group.POST("/queues/:queueName", HandleSendMessage)
	group.GET("/queues/:queueName", HandleConsumeMessage)
	group.POST("/queues/:queueName/purge", HandlePurgeQueue)
}
