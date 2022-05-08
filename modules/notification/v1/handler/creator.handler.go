package handler

import (
	"gin-starter/modules/notification/v1/service"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/resource"
	"gin-starter/response"
)

type NotificationCreatorHandler struct {
	notificationCreator service.NotificationCreatorUseCase
}

func NewNotificationCreatorHandler(
	notificationCreator service.NotificationCreatorUseCase,
) *NotificationCreatorHandler {
	return &NotificationCreatorHandler{
		notificationCreator: notificationCreator,
	}
}

func (cf *NotificationCreatorHandler) CreateNotification(c *gin.Context) {
	var request resource.CreateNotificationRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if err := cf.notificationCreator.InsertNotification(
		c,
		request.UserID,
		request.Title,
		request.Message,
		request.Type,
		request.Extra,
		false,
	); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", nil))
}
