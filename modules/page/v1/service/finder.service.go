package service

import (
	"context"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/page/v1/repository"
)

type PageFinder struct {
	cfg      config.Config
	pageRepo repository.PageRepositoryUseCase
}

type PageFinderUseCase interface {
	GetPages(ctx context.Context) ([]*entity.Page, int64, error)
}

func NewPageFinder(
	cfg config.Config,
	pageRepo repository.PageRepositoryUseCase,
) *PageFinder {
	return &PageFinder{
		cfg:      cfg,
		pageRepo: pageRepo,
	}
}

func (nf *PageFinder) GetPages(ctx context.Context) ([]*entity.Page, int64, error) {
	page, total, err := nf.pageRepo.GetPage(ctx)

	if err != nil {
		return nil, 0, err
	}

	return page, total, nil
}
