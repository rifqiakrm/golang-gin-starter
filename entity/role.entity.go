package entity

import (
	"github.com/google/uuid"
)

const (
	roleTableName = "main.roles"
)

type Role struct {
	ID              uuid.UUID         `json:"id"`
	Name            string            `json:"name"`
	RolePermissions []*RolePermission `foreignKey:"ID" associationForeignKey:"RoleID"`
	Auditable
}

// TableName specifies table name
func (model *Role) TableName() string {
	return roleTableName
}

func NewRole(
	id uuid.UUID,
	name string,
	createdBy string,
) *Role {
	return &Role{
		ID:        id,
		Name:      name,
		Auditable: NewAuditable(createdBy),
	}
}
