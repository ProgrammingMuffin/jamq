package utils

import (
	"jamq-replica/v1/dtos"
	"jamq-replica/v1/services"

	"github.com/labstack/echo/v4"
)

func GetQueueServiceFromRequest(c echo.Context, request interface{}) (services.Queue, error) {
	err := GetRequestBody(c, request)
	if err != nil {
		return nil, echo.NewHTTPError(400, err.Error())
	}
	queueService, err := services.QueueFactory.GetInstance(request.(dtos.CommonQueueRequestData).QueueType)
	if err != nil {
		return nil, echo.ErrBadRequest
	}
	return queueService.(services.Queue), err
}
