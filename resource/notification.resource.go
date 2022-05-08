// Copyright 2021 The starter Authors. All rights reserved.
// This is an API Gateway Resource for starter
// Built with gRPC and Gin Gonic
//
// Notification Resource
package resource

import (
	"github.com/google/uuid"

	"gin-starter/entity"
	"gin-starter/utils"
)

type NotificationRegistrationRequest struct {
	PlayerID string `form:"player_id" json:"player_id" binding:"required"`
	Type     string `form:"type" json:"type" binding:"required"`
}

type CreateNotificationRequest struct {
	UserID  string `form:"user_id" json:"user_id"`
	Title   string `form:"title" json:"title"`
	Message string `form:"message" json:"message"`
	Type    string `form:"type" json:"type"`
	Extra   string `form:"extra" json:"extra"`
}

type GetNotificationsRequest struct {
	Sort   string `form:"sort" json:"sort"`
	Order  string `form:"order" json:"order"`
	Limit  int    `form:"limit,default=10" json:"limit"`
	Offset int    `form:"offset,default=0" json:"offset"`
}

type GetNotificationsResponse struct {
	List  []*Notification `json:"list"`
	Total int64           `json:"total"`
}

type CountUnreadNotificationsResponse struct {
	Total int64 `json:"total"`
}

type ExtraData struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

type Notification struct {
	ID           uuid.UUID  `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Type         string     `json:"type"`
	IsRead       bool       `json:"is_read"`
	Extra        string     `json:"extra"`
	ExtraData    *ExtraData `json:"extra_data"`
	HumanizeTime string     `json:"humanize_time"`
	CreatedAt    string     `json:"created_at"`
	UpdatedAt    string     `json:"updated_at"`
}

func NewNotificationResponse(notification *entity.Notification, extraData *ExtraData) *Notification {
	return &Notification{
		ID:           notification.ID,
		Title:        notification.Title,
		Description:  notification.Description,
		Type:         notification.Type,
		IsRead:       notification.IsRead,
		Extra:        notification.Extra,
		ExtraData:    extraData,
		HumanizeTime: utils.Time(notification.CreatedAt),
		CreatedAt:    notification.CreatedAt.Format(timeFormat),
		UpdatedAt:    notification.UpdatedAt.Format(timeFormat),
	}
}

type CountUnreadNotificationResponse struct {
	Count int64 `json:"count"`
}
