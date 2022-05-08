package entity

import (
	"database/sql"

	"github.com/google/uuid"

	"gin-starter/utils"
)

const (
	notificationTableName        = "main.notifications"
	NotificationTypeAnnouncement = "announcement"
)

type Notification struct {
	ID          uuid.UUID      `gorm:"primaryKey;column:id;" json:"id"`
	UserID      sql.NullString `json:"user_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Type        string         `json:"type"`
	Extra       string         `json:"extra"`
	IsRead      bool           `json:"is_read"`
	Auditable
}

// TableName specifies table name
func (model *Notification) TableName() string {
	return notificationTableName
}

func NewNotification(
	id uuid.UUID,
	userID string,
	title string,
	description string,
	notifType string,
	extra string,
	isRead bool,
	createdBy string,
) *Notification {
	return &Notification{
		ID:          id,
		UserID:      utils.StringToNullString(userID),
		Title:       title,
		Description: description,
		Type:        notifType,
		Extra:       extra,
		IsRead:      isRead,
		Auditable:   NewAuditable(createdBy),
	}
}
