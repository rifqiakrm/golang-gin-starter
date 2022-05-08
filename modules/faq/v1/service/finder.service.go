package service

import (
	"context"
	"gin-starter/config"
	"gin-starter/entity"
	"gin-starter/modules/faq/v1/repository"
)

type FaqFinder struct {
	cfg     config.Config
	faqRepo repository.FaqRepositoryUseCase
}

type FaqFinderUseCase interface {
	GetFaq(ctx context.Context) ([]*entity.Faq, int64, error)
}

func NewFaqFinder(
	cfg config.Config,
	faqRepo repository.FaqRepositoryUseCase,
) *FaqFinder {
	return &FaqFinder{
		cfg:     cfg,
		faqRepo: faqRepo,
	}
}

func (nf *FaqFinder) GetFaq(ctx context.Context) ([]*entity.Faq, int64, error) {
	faq, total, err := nf.faqRepo.GetFaq(ctx)

	if err != nil {
		return nil, 0, err
	}

	return faq, total, nil
}
