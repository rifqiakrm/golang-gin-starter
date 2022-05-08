package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-starter/middleware"
	"gin-starter/modules/notification/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
)

type NotificationUpdaterHandler struct {
	notificationUpdater service.NotificationUpdaterUseCase
}

func NewNotificationUpdaterHandler(
	notificationUpdater service.NotificationUpdaterUseCase,
) *NotificationUpdaterHandler {
	return &NotificationUpdaterHandler{
		notificationUpdater: notificationUpdater,
	}
}

func (cf *NotificationUpdaterHandler) RegisterUnregisterPlayerID(c *gin.Context) {
	var request resource.NotificationRegistrationRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	playerID, err := uuid.Parse(request.PlayerID)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if err := cf.notificationUpdater.RegisterUnregisterPlayerID(
		c,
		middleware.UserID,
		playerID,
		request.Type,
	); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}

func (cf *NotificationUpdaterHandler) UpdateReadNotification(c *gin.Context) {
	if err := cf.notificationUpdater.UpdateReadNotification(
		c,
		middleware.UserID,
	); err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}
