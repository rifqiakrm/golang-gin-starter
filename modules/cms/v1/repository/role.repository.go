package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-starter/entity"
)

type RoleRepository struct {
	db *gorm.DB
}

type RoleRepositoryUseCase interface {
	Create(ctx context.Context, role *entity.Role) error
	FindAll(ctx context.Context) ([]*entity.Role, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (nc *RoleRepository) Create(ctx context.Context, role *entity.Role) error {
	if err := nc.db.
		WithContext(ctx).
		Model(&entity.Role{}).
		Create(role).
		Error; err != nil {
		return errors.Wrap(err, "[RoleRepository-CreateNews] error while creating user")
	}

	return nil
}

func (nc *RoleRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	category := &entity.Role{}

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.Role{}).
		Preload("RolePermissions").
		Preload("RolePermissions.Permission").
		Where("id = ?", id).
		First(&category).
		Error; err != nil {
		return nil, errors.Wrap(err, "[NewsRepositoryRepository-FindByID] error while getting category category")
	}

	return category, nil
}

func (nc *RoleRepository) FindAll(ctx context.Context) ([]*entity.Role, error) {
	role := make([]*entity.Role, 0)

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.Role{}).
		Preload("RolePermissions").
		Preload("RolePermissions.Permission").
		Find(&role).
		Error; err != nil {
		return nil, errors.Wrap(err, "[RoleRepository-GetNewsCategories] error while getting news category")
	}

	return role, nil
}

func (nc *RoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := nc.db.WithContext(ctx).
		Model(&entity.Role{}).
		Where(`id = ?`, id).
		Updates(
			map[string]interface{}{
				"updated_at": time.Now(),
				"deleted_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
	}

	return nil
}
