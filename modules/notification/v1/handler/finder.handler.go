package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-starter/entity"
	"gin-starter/middleware"
	service2 "gin-starter/modules/notification/v1/service"
	"gin-starter/resource"
	"gin-starter/response"
)

type NotificationFinderHandler struct {
	notificationFinder  service2.NotificationFinderUseCase
	notificationUpdater service2.NotificationUpdaterUseCase
}

func NewNotificationFinderHandler(
	notificationFinder service2.NotificationFinderUseCase,
	notificationUpdater service2.NotificationUpdaterUseCase,
) *NotificationFinderHandler {
	return &NotificationFinderHandler{
		notificationFinder:  notificationFinder,
		notificationUpdater: notificationUpdater,
	}
}

func (cf *NotificationFinderHandler) GetNotification(c *gin.Context) {
	var request resource.GetNotificationsRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	notifications, total, err := cf.notificationFinder.GetNotification(
		c,
		middleware.UserID,
		request.Sort,
		request.Order,
		request.Limit,
		request.Offset,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	res := make([]*resource.Notification, 0)
	var extraData *resource.ExtraData
	for _, n := range notifications {
		if n.Extra != "" {
			switch n.Type {
			case entity.NotificationTypeAnnouncement:
			default:
				extraData = &resource.ExtraData{}
			}
		}
		res = append(res, resource.NewNotificationResponse(n, extraData))
	}

	if err := cf.notificationUpdater.UpdateReadNotification(c, middleware.UserID); err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", &resource.GetNotificationsResponse{
		List:  res,
		Total: total,
	}))
}

func (cf *NotificationFinderHandler) CountUnreadNotifications(c *gin.Context) {
	total, err := cf.notificationFinder.CountUnreadNotifications(
		c,
		middleware.UserID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.ErrorApiResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessApiResponseList(http.StatusOK, "success", &resource.CountUnreadNotificationsResponse{
		Total: total,
	}))
}
