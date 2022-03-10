package v1

import (
	"jamq-replica/utils"
	"jamq-replica/v1/dtos"
	"jamq-replica/v1/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleCreateQueue(c echo.Context) error {
	request := &dtos.CreateQueueRequest{}
	queueService, err := utils.GetQueueServiceFromRequest(c, request)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	err = queueService.(services.Queue).CreateQueue(request.QueueName)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	response := &dtos.CreateQueueResponse{Response: "Successfully created queue"}
	return c.JSON(http.StatusOK, response)
}

func HandlePurgeQueue(c echo.Context) error {
	request := &dtos.PurgeQueueRequest{}
	queueService, err := utils.GetQueueServiceFromRequest(c, request)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	err = queueService.PurgeQueue(request.QueueName)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	response := &dtos.PurgeQueueResponse{Response: "Successfully purged queue"}
	return c.JSON(http.StatusOK, response)
}

func HandleSendMessage(c echo.Context) error {
	request := &dtos.SendMessageRequest{}
	queueService, err := utils.GetQueueServiceFromRequest(c, request)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	err = queueService.SendMessage(request.QueueName, request.Message, request.Timestamp)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	response := &dtos.SendMessageResponse{Response: "Successfully added message to queue"}
	return c.JSON(http.StatusOK, response)
}

func HandleConsumeMessage(c echo.Context) error {
	request := &dtos.ReceiveMessageRequest{}
	queueService, err := utils.GetQueueServiceFromRequest(c, request)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	node, err := queueService.ConsumeMessage(request.QueueName)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	response := &dtos.ConsumeMessageResponse{Message: node.Message, Timestamp: node.Timestamp}
	return c.JSON(http.StatusOK, response)
}
