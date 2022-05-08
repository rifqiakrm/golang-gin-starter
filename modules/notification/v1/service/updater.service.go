package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/rifqiakrm/onesignal-go-lib"

	"gin-starter/config"
	"gin-starter/modules/notification/v1/repository"
)

type NotificationUpdater struct {
	cfg              config.Config
	notificationRepo repository.NotificationRepositoryUseCase
}

type NotificationUpdaterUseCase interface {
	RegisterUnregisterPlayerID(ctx context.Context, userID uuid.UUID, playerID uuid.UUID, typeReg string) error
	UpdateReadNotification(ctx context.Context, id uuid.UUID) error
}

func NewNotificationUpdater(
	cfg config.Config,
	notificationRepo repository.NotificationRepositoryUseCase,
) *NotificationUpdater {
	return &NotificationUpdater{
		cfg:              cfg,
		notificationRepo: notificationRepo,
	}
}

func (nu *NotificationUpdater) RegisterUnregisterPlayerID(ctx context.Context, userID uuid.UUID, playerID uuid.UUID, typeReg string) error {
	var tags map[string]string

	if typeReg == "reg" {
		tags = map[string]string{
			"uuid":      userID.String(),
			"logged_in": "true",
		}
	} else {
		tags = map[string]string{
			"uuid":      "",
			"logged_in": "false",
		}
	}

	client := onesignal.NewClient(nil)

	player := &onesignal.PlayerRequest{
		AppID: nu.cfg.OneSignal.AppID,
		Tags:  tags,
	}

	_, _, err := client.Players.Update(playerID.String(), player)

	if err != nil {
		return err
	}

	return nil
}

func (nu *NotificationUpdater) UpdateReadNotification(ctx context.Context, id uuid.UUID) error {
	if err := nu.notificationRepo.UpdateReadNotification(ctx, id); err != nil {
		return err
	}

	return nil
}
