package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-starter/entity"
)

type NotificationRepository struct {
	db *gorm.DB
}

type NotificationRepositoryUseCase interface {
	GetNotification(ctx context.Context, id uuid.UUID, sort, order string, limit, offset int) ([]*entity.Notification, int64, error)
	Create(ctx context.Context, notification *entity.Notification) error
	CountUnreadNotification(ctx context.Context, id uuid.UUID) (int64, error)
	UpdateReadNotification(ctx context.Context, id uuid.UUID) error
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db}
}

func (nr *NotificationRepository) GetNotification(ctx context.Context, id uuid.UUID, sort, order string, limit, offset int) ([]*entity.Notification, int64, error) {
	notifications := make([]*entity.Notification, 0)
	var total int64
	var gormDB = nr.db.
		WithContext(ctx).
		Model(&entity.Notification{}).
		Where("user_id = ?", id).
		Or("user_id is null")

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(offset)

	if err := gormDB.Find(&notifications).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[NotificationRepository-GetNotification] error while retrieving notifications data")
	}

	return notifications, total, nil
}

func (nr *NotificationRepository) Create(ctx context.Context, notification *entity.Notification) error {
	if err := nr.db.
		WithContext(ctx).
		Model(&entity.Notification{}).
		Create(notification).
		Error; err != nil {
		return errors.Wrap(err, "[NewsCategoryRepository-CreateNews] error while creating user")
	}

	return nil
}

func (nr *NotificationRepository) CountUnreadNotification(ctx context.Context, id uuid.UUID) (int64, error) {
	var total int64

	nr.db.
		WithContext(ctx).
		Model(&entity.Notification{}).
		Where("user_id = ?", id).
		Or("user_id is NULL").
		Where("is_read = false").
		Count(&total)

	return total, nil
}

func (nr *NotificationRepository) UpdateReadNotification(ctx context.Context, id uuid.UUID) error {
	if err := nr.db.WithContext(ctx).
		Model(&entity.Notification{}).
		Where(`user_id = ?`, id).
		Or("user_id is NULL").
		Updates(
			map[string]interface{}{
				"is_read":    true,
				"updated_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
	}

	return nil
}
