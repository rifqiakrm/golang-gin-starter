package service

import (
	"context"
	"gin-starter/modules/cms/v1/repository"

	"github.com/google/uuid"

	"gin-starter/config"
	"gin-starter/entity"
)

type CMSFinder struct {
	cfg      config.Config
	userRepo repository.UserRepositoryUseCase
	roleRepo repository.RoleRepositoryUseCase
	faqRepo  repository.FaqRepositoryUseCase
	pageRepo repository.PageRepositoryUseCase
}

type CMSFinderUseCase interface {
	GetUsers(ctx context.Context, query, order, sort string, limit, offset int) ([]*entity.User, int64, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetAdminUsers(ctx context.Context, query, order, sort string, limit, offset int) ([]*entity.User, int64, error)
	GetAdminUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetRoles(ctx context.Context) ([]*entity.Role, error)
	GetFaqByID(ctx context.Context, id uuid.UUID) (*entity.Faq, error)
	GetFaqs(ctx context.Context) ([]*entity.Faq, error)
	GetPageByID(ctx context.Context, id uuid.UUID) (*entity.Page, error)
	GetPages(ctx context.Context) ([]*entity.Page, error)
}

func NewCMSFinder(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
	faqRepo repository.FaqRepositoryUseCase,
	pageRepo repository.PageRepositoryUseCase,
) *CMSFinder {
	return &CMSFinder{
		cfg:      cfg,
		userRepo: userRepo,
		roleRepo: roleRepo,
		faqRepo:  faqRepo,
		pageRepo: pageRepo,
	}
}

func (cf *CMSFinder) GetUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error) {
	users, total, err := cf.userRepo.GetUsers(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (cf *CMSFinder) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	users, err := cf.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (cf *CMSFinder) GetAdminUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error) {
	users, total, err := cf.userRepo.GetAdminUsers(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (cf *CMSFinder) GetAdminUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	users, err := cf.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (cf *CMSFinder) GetRoles(ctx context.Context) ([]*entity.Role, error) {
	roles, err := cf.roleRepo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (cf *CMSFinder) GetFaqByID(ctx context.Context, id uuid.UUID) (*entity.Faq, error) {
	faq, err := cf.faqRepo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return faq, nil
}

func (cf *CMSFinder) GetFaqs(ctx context.Context) ([]*entity.Faq, error) {
	faqs, err := cf.faqRepo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return faqs, nil
}

func (cf *CMSFinder) GetPageByID(ctx context.Context, id uuid.UUID) (*entity.Page, error) {
	page, err := cf.pageRepo.FindByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return page, nil
}

func (cf *CMSFinder) GetPages(ctx context.Context) ([]*entity.Page, error) {
	pages, err := cf.pageRepo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return pages, nil
}
