package resource

import (
	"github.com/google/uuid"

	"gin-starter/entity"
)

type GetRoleResponse struct {
	List  []*Role `json:"list"`
	Total int64   `json:"total"`
}

type GetRoleByID struct {
	ID string `uri:"id" binding:"required"`
}

type CreateRoleRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type UpdateRoleRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type DeleteRoleRequest struct {
	ID string `uri:"id" binding:"required"`
}

type Role struct {
	ID         uuid.UUID     `json:"id"`
	Name       string        `json:"name"`
	Permission []*Permission `json:"permissions"`
}

func NewRoleResponse(role *entity.Role) *Role {
	permissions := make([]*Permission, 0)

	for _, v := range role.RolePermissions {
		permissions = append(permissions, &Permission{
			ID:   v.Permission.ID.String(),
			Name: v.Permission.Name,
		})
	}

	return &Role{
		ID:         role.ID,
		Name:       role.Name,
		Permission: permissions,
	}
}
