package v1

import "github.com/labstack/echo/v4"

func HandleV1Routes(group *echo.Group) {
	group.POST("/queues/:queueName", HandleCreateQueue)
	group.PUT("/queues/:queueName", HandleSendMessage)
	group.GET("/queues/:queueName", HandleConsumeMessage)
	group.PATCH("/queues/:queueName", HandlePurgeQueue)
}
