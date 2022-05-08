package service

import (
	"context"
	"time"

	"github.com/google/uuid"

	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/cms/v1/repository"
	notificationService "gin-starter/modules/notification/v1/service"
	"gin-starter/utils"
)

type CMSCreator struct {
	cfg          config.Config
	userRepo     repository.UserRepositoryUseCase
	userRoleRepo repository.UserRoleRepositoryUseCase
	notifCreator notificationService.NotificationCreatorUseCase
	faqRepo      repository.FaqRepositoryUseCase
	pageRepo     repository.PageRepositoryUseCase
}

type CMSCreatorUseCase interface {
	CreateUser(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time) (*entity.User, error)
	CreateAdmin(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time, roleID uuid.UUID) (*entity.User, error)
	CreateFaq(ctx context.Context, question string, answer string) (*entity.Faq, error)
	CreatePage(ctx context.Context, title string, content string) (*entity.Page, error)
}

func NewCMSCreator(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	notifCreator notificationService.NotificationCreatorUseCase,
	faqRepo repository.FaqRepositoryUseCase,
	pageRepo repository.PageRepositoryUseCase,
) *CMSCreator {
	return &CMSCreator{
		cfg:          cfg,
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		notifCreator: notifCreator,
		faqRepo:      faqRepo,
		pageRepo:     pageRepo,
	}
}

func (cc *CMSCreator) CreateUser(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time) (*entity.User, error) {
	user := entity.NewUser(
		uuid.New(),
		name,
		email,
		password,
		utils.TimeToNullTime(dob),
		photo,
		phoneNumber,
		"system",
	)

	if err := cc.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (cc *CMSCreator) CreateAdmin(ctx context.Context, name, email, password, phoneNumber, photo string, dob time.Time, roleID uuid.UUID) (*entity.User, error) {
	userID := uuid.New()
	user := entity.NewUser(
		userID,
		name,
		email,
		password,
		utils.TimeToNullTime(dob),
		photo,
		phoneNumber,
		"system",
	)

	if err := cc.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	userRole := entity.NewUserRole(userID, roleID, "system")

	if err := cc.userRoleRepo.CreateOrUpdate(ctx, userRole); err != nil {
		return nil, err
	}

	return user, nil
}

func (cc *CMSCreator) CreateFaq(ctx context.Context, question string, answer string) (*entity.Faq, error) {
	faq := entity.NewFaq(
		uuid.New(),
		question,
		answer,
		"system",
	)

	if err := cc.faqRepo.Create(ctx, faq); err != nil {
		return nil, err
	}

	return faq, nil
}

func (cc *CMSCreator) CreatePage(ctx context.Context, title string, content string) (*entity.Page, error) {
	page := entity.NewPage(
		uuid.New(),
		title,
		content,
		"system",
	)

	if err := cc.pageRepo.Create(ctx, page); err != nil {
		return nil, err
	}

	return page, nil
}
