package repository

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gin-starter/entity"
	"gin-starter/utils"
)

type AuthRepository struct {
	db *gorm.DB
}

type AuthRepositoryUseCase interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetAdminByEmail(ctx context.Context, email string) (*entity.User, error)
	UpdateOTP(ctx context.Context, user *entity.User, otp string) error
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (ar *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
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

func (ar *AuthRepository) UpdateOTP(ctx context.Context, user *entity.User, otp string) error {
	oldTime := user.UpdatedAt
	user.UpdatedAt = time.Now()
	if err := ar.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.User)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, user.ID).Error; err != nil {
				return errors.Wrap(err, "[UserRepository-ChangePassword] error when updating data")
			}
			if err := tx.Model(&entity.User{}).
				Where(`id = ?`, user.ID).Update("otp", utils.StringToNullString(otp)).Error; err != nil {
				return errors.Wrap(err, "[UserRepository-Update] error when updating data User b")
			}
			return nil
		}); err != nil {
		user.UpdatedAt = oldTime
	}
	return nil
}

func (ar *AuthRepository) GetAdminByEmail(ctx context.Context, email string) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Joins("inner join user_roles on users.id=user_roles.user_id").
		Where("email = ?", email).
		Find(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetAdminByEmail] email not found")
	}

	return result, nil
}
