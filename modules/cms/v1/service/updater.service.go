package service

import (
	"context"
	"gin-starter/modules/cms/v1/repository"
	notificationService "gin-starter/modules/notification/v1/service"

	"github.com/google/uuid"

	"gin-starter/config"
	"gin-starter/entity"
)

type CMSUpdater struct {
	cfg          config.Config
	userRepo     repository.UserRepositoryUseCase
	userRoleRepo repository.UserRoleRepositoryUseCase
	notifCreator notificationService.NotificationCreatorUseCase
	faqRepo      repository.FaqRepositoryUseCase
	pageRepo     repository.PageRepositoryUseCase
}

type CMSUpdaterUseCase interface {
	ActivateDeactivateUser(ctx context.Context, id uuid.UUID) error
	UpdateAdmin(ctx context.Context, user *entity.User, roleID uuid.UUID) error
	UpdateFaq(ctx context.Context, faq *entity.Faq) error
	UpdatePage(ctx context.Context, page *entity.Page) error
}

func NewCMSUpdater(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	notifCreator notificationService.NotificationCreatorUseCase,
	faqRepo repository.FaqRepositoryUseCase,
	pageRepo repository.PageRepositoryUseCase,
) *CMSUpdater {
	return &CMSUpdater{
		cfg:          cfg,
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		notifCreator: notifCreator,
		faqRepo:      faqRepo,
		pageRepo:     pageRepo,
	}
}

func (cu *CMSUpdater) ActivateDeactivateUser(ctx context.Context, id uuid.UUID) error {
	user, err := cu.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return err
	}

	if user == nil {
		return entity.ErrUserNotFound.Error
	}

	if user.Status == "DEACTIVATED" {
		if err := cu.userRepo.UpdateUserStatus(ctx, id, "ACTIVATED"); err != nil {
			return err
		}
	} else if user.Status == "ACTIVATED" {
		if err := cu.userRepo.UpdateUserStatus(ctx, id, "DEACTIVATED"); err != nil {
			return err
		}
	}

	return nil
}

func (cu *CMSUpdater) UpdateAdmin(ctx context.Context, user *entity.User, roleID uuid.UUID) error {
	if err := cu.userRepo.UpdateUser(ctx, user); err != nil {
		return err
	}

	userRole, err := cu.userRoleRepo.FindByUserID(ctx, user.ID)

	if err != nil {
		return err
	}

	userRole.RoleID = roleID

	if err := cu.userRoleRepo.Update(ctx, userRole); err != nil {
		return nil
	}

	return nil
}

func (cu *CMSUpdater) UpdateFaq(ctx context.Context, faq *entity.Faq) error {
	if err := cu.faqRepo.Update(ctx, faq); err != nil {
		return err
	}

	return nil
}

func (cu *CMSUpdater) UpdatePage(ctx context.Context, page *entity.Page) error {
	if err := cu.pageRepo.Update(ctx, page); err != nil {
		return err
	}

	return nil
}
