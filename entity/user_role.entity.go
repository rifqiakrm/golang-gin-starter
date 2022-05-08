package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	userRoleTableName = "main.user_roles"
)

type UserRole struct {
	UserID uuid.UUID `json:"user_id"`
	RoleID uuid.UUID `json:"role_id"`
	Role   *Role     `foreignKey:"RoleID"`
	Auditable
}

// TableName specifies table name
func (model *UserRole) TableName() string {
	return userRoleTableName
}

func NewUserRole(
	userID uuid.UUID,
	roleID uuid.UUID,
	createdBy string,
) *UserRole {
	return &UserRole{
		UserID:    userID,
		RoleID:    roleID,
		Auditable: NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *UserRole) MapUpdateFrom(from *UserRole) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"role_id":    model.RoleID,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.RoleID != from.RoleID {
		mapped["role_id"] = from.RoleID
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
