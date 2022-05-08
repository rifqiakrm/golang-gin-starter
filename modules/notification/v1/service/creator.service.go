package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rifqiakrm/onesignal-go-lib"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/notification/v1/repository"
)

type NotificationCreator struct {
	cfg              config.Config
	notificationRepo repository.NotificationRepositoryUseCase
}

type NotificationCreatorUseCase interface {
	InsertNotification(ctx context.Context, userID string, title, message, notifType, extra string, isRead bool) error
}

func NewNotificationCreator(
	cfg config.Config,
	notificationRepo repository.NotificationRepositoryUseCase,
) *NotificationCreator {
	return &NotificationCreator{
		cfg:              cfg,
		notificationRepo: notificationRepo,
	}
}

func (nc *NotificationCreator) InsertNotification(ctx context.Context, userID string, title, message, notifType, extra string, isRead bool) error {
	client := onesignal.NewClient(nil)
	client.AppKey = nc.cfg.OneSignal.AppKey

	notificationReq := &onesignal.NotificationRequest{
		AppID: nc.cfg.OneSignal.AppID,
		Contents: map[string]string{
			"en": message,
		},
		Headings: map[string]string{
			"en": title,
		},
		Tags: []map[string]string{
			{
				"key":      "logged_in",
				"relation": "=",
				"value":    "true",
			},
		},
		Data: nil,
	}

	_, _, err := client.Notifications.Create(notificationReq)

	if err != nil {
		return err
	}

	notification := entity.NewNotification(
		uuid.New(),
		userID,
		title,
		message,
		notifType,
		extra,
		isRead,
		"system",
	)

	if err := nc.notificationRepo.Create(ctx, notification); err != nil {
		return err
	}

	return nil
}
