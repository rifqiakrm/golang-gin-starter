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

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryUseCase interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetUserByForgotPasswordToken(ctx context.Context, token string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	ChangePassword(ctx context.Context, user *entity.User, newPassword string) error
	UpdateOTP(ctx context.Context, user *entity.User, otp string) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ar *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Where("email = ?", email).
		Find(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByEmail] email not found")
	}

	return result, nil
}

func (ar *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByID] user not found")
	}

	return result, nil
}

func (ar *UserRepository) GetUserByForgotPasswordToken(ctx context.Context, token string) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Where("forgot_password_token = ?", token).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByID] user not found")
	}

	return result, nil
}

func (ar *UserRepository) UpdateOTP(ctx context.Context, user *entity.User, otp string) error {
	if err := ar.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, user.ID).
		Updates(
			map[string]interface{}{
				"otp":        otp,
				"updated_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-Update] error when updating user data")
	}
	return nil
}

func (ar *UserRepository) Update(ctx context.Context, user *entity.User) error {
	oldTime := user.UpdatedAt
	user.UpdatedAt = time.Now()
	if err := ar.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.User)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, user.ID).Error; err != nil {
				log.Println("[GamPTKRepository - Update]", err)
				return err
			}
			if err := tx.Model(&entity.User{}).
				Where(`id`, user.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(user)).Error; err != nil {
				log.Println("[GamPTKRepository - Update]", err)
				return err
			}
			return nil
		}); err != nil {
		user.UpdatedAt = oldTime
	}
	return nil
}

func (ar *UserRepository) ChangePassword(ctx context.Context, user *entity.User, newPassword string) error {
	if err := ar.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, user.ID).
		Updates(
			map[string]interface{}{
				"password":   newPassword,
				"updated_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-Update] error when updating user data")
	}

	return nil
}
