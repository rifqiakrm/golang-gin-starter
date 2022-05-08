package service

import (
	"context"

	"github.com/google/uuid"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/user/v1/repository"
)

type UserFinder struct {
	cfg      config.Config
	userRepo repository.UserRepositoryUseCase
}

type UserFinderUseCase interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByForgotPasswordToken(ctx context.Context, token string) (*entity.User, error)
}

func NewUserFinder(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
) *UserFinder {
	return &UserFinder{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (as *UserFinder) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := as.userRepo.GetUserByEmail(ctx, email)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (as *UserFinder) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := as.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (as *UserFinder) GetUserByForgotPasswordToken(ctx context.Context, token string) (*entity.User, error) {
	user, err := as.userRepo.GetUserByForgotPasswordToken(ctx, token)

	if err != nil {
		return user, err
	}

	return user, nil
}
