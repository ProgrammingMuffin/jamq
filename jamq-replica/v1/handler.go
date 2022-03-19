package v1

import (
	"fmt"
	"jamq-replica/utils"
	"jamq-replica/v1/dtos"
	"jamq-replica/v1/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleCreateQueue(c echo.Context) error {
	fmt.Println("handle create queue")
	queueService, err := utils.GetQueueServiceFromRequest(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	err = queueService.(services.Queue).CreateQueue(c.Param("queueName"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	response := &dtos.CreateQueueResponse{Response: "Successfully created queue"}
	return c.JSON(http.StatusOK, response)
}

func HandlePurgeQueue(c echo.Context) error {
	queueService, err := utils.GetQueueServiceFromRequest(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	err = queueService.PurgeQueue(c.Param("queueName"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	response := &dtos.PurgeQueueResponse{Response: "Successfully purged queue"}
	return c.JSON(http.StatusOK, response)
}

func HandleSendMessage(c echo.Context) error {
	request := &dtos.SendMessageRequest{}
	queueService, err := utils.GetQueueServiceFromRequest(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	err = queueService.SendMessage(c.Param("queueName"), request.Message, request.Timestamp)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	response := &dtos.SendMessageResponse{Response: "Successfully added message to queue"}
	return c.JSON(http.StatusOK, response)
}

func HandleConsumeMessage(c echo.Context) error {
	queueService, err := utils.GetQueueServiceFromRequest(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	node, err := queueService.ConsumeMessage(c.Param("queueName"))
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}
	response := &dtos.ConsumeMessageResponse{Message: node.Message, Timestamp: node.Timestamp}
	return c.JSON(http.StatusOK, response)
}
