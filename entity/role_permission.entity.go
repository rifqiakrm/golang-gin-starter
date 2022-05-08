package entity

import (
	"github.com/google/uuid"
)

const (
	rolePermissionTableName = "main.role_permissions"
)

type RolePermission struct {
	ID           uuid.UUID   `json:"id"`
	RoleID       uuid.UUID   `json:"role_id"`
	PermissionID uuid.UUID   `json:"permission_id"`
	Permission   *Permission `gorm:"ForeignKey:PermissionID;AssociationForeignKey:ID` //nolint
	Auditable
}

// TableName specifies table name
func (model *RolePermission) TableName() string {
	return rolePermissionTableName
}

func NewRolePermission(
	id uuid.UUID,
	roleID uuid.UUID,
	permissionID uuid.UUID,
	createdBy string,
) *RolePermission {
	return &RolePermission{
		ID:           id,
		RoleID:       roleID,
		PermissionID: permissionID,
		Auditable:    NewAuditable(createdBy),
	}
}
