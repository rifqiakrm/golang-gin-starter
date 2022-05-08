package repository

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gin-starter/entity"
)

type UserRoleRepository struct {
	db *gorm.DB
}

type UserRoleRepositoryUseCase interface {
	CreateOrUpdate(ctx context.Context, userRole *entity.UserRole) error
	FindByUserID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error)
	Update(ctx context.Context, userRole *entity.UserRole) error
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewUserRoleRepository(db *gorm.DB) *UserRoleRepository {
	return &UserRoleRepository{db}
}

func (nc *UserRoleRepository) CreateOrUpdate(ctx context.Context, userRole *entity.UserRole) error {
	var find *entity.UserRole

	findUser := nc.db.
		Where("user_id = ?", userRole.UserID).
		First(&find)

	if err := findUser.Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	if findUser.RowsAffected > 0 {
		if err := nc.db.Model(&entity.UserRole{}).
			Where("user_id = ?", userRole.UserID).
			UpdateColumns(map[string]interface{}{
				"role_id": userRole.RoleID,
			}).
			Error; err != nil {
			return err
		}

		return nil
	}

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.UserRole{}).
		Create(userRole).
		Error; err != nil {
		return errors.Wrap(err, "[UserRoleRepository-CreateNews] error while creating user")
	}

	return nil
}

func (nc *UserRoleRepository) FindByUserID(ctx context.Context, id uuid.UUID) (*entity.UserRole, error) {
	category := &entity.UserRole{}

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.UserRole{}).
		Where("user_id = ?", id).
		First(&category).
		Error; err != nil {
		return nil, errors.Wrap(err, "[NewsRepositoryRepository-FindByID] error while getting category category")
	}

	return category, nil
}

func (nc *UserRoleRepository) Update(ctx context.Context, userRole *entity.UserRole) error {
	oldTime := userRole.UpdatedAt
	userRole.UpdatedAt = time.Now()
	if err := nc.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.UserRole)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("user_id = ?", userRole.UserID).
				Find(&sourceModel).Error; err != nil {
				log.Println("[GamPTKRepository - Update]", err)
				return err
			}
			if err := tx.Model(&entity.UserRole{}).
				Where(`user_id`, userRole.UserID).
				UpdateColumns(sourceModel.MapUpdateFrom(userRole)).Error; err != nil {
				log.Println("[GamPTKRepository - Update]", err)
				return err
			}
			return nil
		}); err != nil {
		userRole.UpdatedAt = oldTime
	}
	return nil
}

func (nc *UserRoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := nc.db.WithContext(ctx).
		Model(&entity.UserRole{}).
		Where(`user_id = ?`, id).
		Updates(
			map[string]interface{}{
				"updated_at": time.Now(),
				"deleted_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
	}

	return nil
}
