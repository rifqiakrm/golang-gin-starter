package service

import (
	"context"
	"gin-starter/modules/cms/v1/repository"

	"github.com/google/uuid"

	"gin-starter/config"
)

type CMSDeleter struct {
	cfg      config.Config
	userRepo repository.UserRepositoryUseCase
	faqRepo  repository.FaqRepositoryUseCase
	pageRepo repository.PageRepositoryUseCase
}

type CMSDeleterUseCase interface {
	DeleteAdmin(ctx context.Context, id uuid.UUID) error
	DeleteFaq(ctx context.Context, id uuid.UUID) error
	DeletePage(ctx context.Context, id uuid.UUID) error
}

func NewCMSDeleter(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	faqRepo repository.FaqRepositoryUseCase,
	pageRepo repository.PageRepositoryUseCase,
) *CMSDeleter {
	return &CMSDeleter{
		cfg:      cfg,
		userRepo: userRepo,
		faqRepo:  faqRepo,
		pageRepo: pageRepo,
	}
}

func (cd *CMSDeleter) DeleteAdmin(ctx context.Context, id uuid.UUID) error {
	if err := cd.userRepo.DeleteAdmin(ctx, id); err != nil {
		return err
	}

	return nil
}

func (cd *CMSDeleter) DeleteFaq(ctx context.Context, id uuid.UUID) error {
	if err := cd.faqRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func (cd *CMSDeleter) DeletePage(ctx context.Context, id uuid.UUID) error {
	if err := cd.pageRepo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
