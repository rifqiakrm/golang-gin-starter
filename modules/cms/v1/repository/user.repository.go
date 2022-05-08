package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gin-starter/entity"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error)
	GetAdminUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	UpdateUserStatus(ctx context.Context, id uuid.UUID, status string) error
	DeleteAdmin(ctx context.Context, id uuid.UUID) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := ur.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Create(user).
		Error; err != nil {
		return errors.Wrap(err, "[UserRepository-CreateUser] error while creating user")
	}

	return nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	result := new(entity.User)

	if err := ur.db.
		WithContext(ctx).
		Preload("UserRole").
		Preload("UserRole.Role.RolePermissions").
		Preload("UserRole.Role.RolePermissions.Permission").
		Where("id = ?", id).
		Find(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByID] user not found")
	}

	return result, nil
}

func (ur *UserRepository) GetUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error) {
	var user []*entity.User
	var total int64
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Joins("left join user_roles on users.id=user_roles.user_id").
		Where("user_roles.user_id is null")

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(offset)

	if query != "" {
		gormDB = gormDB.
			Where("name ILIKE ?", "%"+query+"%").
			Or("email ILIKE ?", "%"+query+"%").
			Or("phone_number ILIKE ?", "%"+query+"%")
	}

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := gormDB.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[UserRepository-GetAdminUsers] error when looking up all user")
	}

	return user, total, nil
}

func (ur *UserRepository) GetAdminUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error) {
	var user []*entity.User
	var total int64
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Preload("UserRole").
		Preload("UserRole.Role.RolePermissions").
		Preload("UserRole.Role.RolePermissions.Permission").
		Joins("inner join user_roles on users.id=user_roles.user_id")

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(offset)

	if query != "" {
		gormDB = gormDB.
			Where("name ILIKE ?", "%"+query+"%").
			Or("email ILIKE ?", "%"+query+"%").
			Or("phone_number ILIKE ?", "%"+query+"%")
	}

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", sort, order))

	if err := gormDB.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[UserRepository-GetAdminUsers] error when looking up all user")
	}

	return user, total, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	oldTime := user.UpdatedAt
	user.UpdatedAt = time.Now()
	if err := ur.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.User)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, user.ID).Error; err != nil {
				log.Println("[UserRepository-UpdateUser]", err)
				return err
			}
			if err := tx.Model(&entity.User{}).
				Where(`id`, user.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(user)).Error; err != nil {
				log.Println("[UserRepository-UpdateUser]", err)
				return err
			}
			return nil
		}); err != nil {
		user.UpdatedAt = oldTime
	}
	return nil
}

func (ur *UserRepository) UpdateUserStatus(ctx context.Context, id uuid.UUID, status string) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, id).
		Updates(
			map[string]interface{}{
				"status":     status,
				"updated_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
	}

	return nil
}

func (ur *UserRepository) DeleteAdmin(ctx context.Context, id uuid.UUID) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, id).
		Delete(&entity.User{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeleteAdmin] error when updating user data")
	}

	return nil
}
