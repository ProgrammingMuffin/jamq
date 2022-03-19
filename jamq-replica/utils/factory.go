package utils

import (
	"jamq-replica/constants"
	"jamq-replica/v1/services"

	"github.com/labstack/echo/v4"
)

func GetQueueServiceFromRequest(c echo.Context) (services.Queue, error) {
	queueType := c.QueryParam(constants.GetQueryParamKeys().QueueType)
	if queueType == "" {
		return nil, echo.ErrBadRequest
	}
	queueService, err := services.QueueFactory.GetInstance(queueType)
	if err != nil {
		return nil, echo.ErrBadRequest
	}
	return queueService.(services.Queue), err
}
