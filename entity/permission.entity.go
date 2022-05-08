package entity

import (
	"github.com/google/uuid"
)

const (
	permissionTableName = "main.permissions"
)

type Permission struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Auditable
}

// TableName specifies table name
func (model *Permission) TableName() string {
	return permissionTableName
}

func NewPermission(
	id uuid.UUID,
	name string,
	createdBy string,
) *Permission {
	return &Permission{
		ID:        id,
		Name:      name,
		Auditable: NewAuditable(createdBy),
	}
}
