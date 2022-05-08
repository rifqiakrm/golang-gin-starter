package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-starter/entity"
)

type FaqRepository struct {
	db *gorm.DB
}

type FaqRepositoryUseCase interface {
	GetFaq(ctx context.Context) ([]*entity.Faq, int64, error)
}

func NewFaqRepository(db *gorm.DB) *FaqRepository {
	return &FaqRepository{db}
}

func (nr *FaqRepository) GetFaq(ctx context.Context) ([]*entity.Faq, int64, error) {
	faq := make([]*entity.Faq, 0)
	var total int64
	var gormDB = nr.db.
		WithContext(ctx).
		Model(&entity.Faq{})

	gormDB = gormDB.Order("created_at ASC")

	if err := gormDB.Find(&faq).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[FaqRepository-GetFaq] error while retrieving faq data")
	}

	return faq, total, nil
}
