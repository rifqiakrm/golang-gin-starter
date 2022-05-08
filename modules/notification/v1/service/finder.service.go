package service

import (
	"context"

	"github.com/google/uuid"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/notification/v1/repository"
)

type NotificationFinder struct {
	cfg              config.Config
	notificationRepo repository.NotificationRepositoryUseCase
}

type NotificationFinderUseCase interface {
	GetNotification(ctx context.Context, id uuid.UUID, sort, order string, limit, offset int) ([]*entity.Notification, int64, error)
	CountUnreadNotifications(ctx context.Context, id uuid.UUID) (int64, error)
}

func NewNotificationFinder(
	cfg config.Config,
	notificationRepo repository.NotificationRepositoryUseCase,
) *NotificationFinder {
	return &NotificationFinder{
		cfg:              cfg,
		notificationRepo: notificationRepo,
	}
}

func (nf *NotificationFinder) GetNotification(ctx context.Context, id uuid.UUID, sort, order string, limit, offset int) ([]*entity.Notification, int64, error) {
	notification, total, err := nf.notificationRepo.GetNotification(
		ctx,
		id,
		sort,
		order,
		limit,
		offset,
	)

	if err != nil {
		return nil, 0, err
	}

	return notification, total, nil
}

func (nf *NotificationFinder) CountUnreadNotifications(ctx context.Context, id uuid.UUID) (int64, error) {
	total, err := nf.notificationRepo.CountUnreadNotification(ctx, id)

	if err != nil {
		return 0, err
	}

	return total, nil
}
