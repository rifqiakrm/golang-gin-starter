package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"gin-starter/entity"
)

type PageRepository struct {
	db *gorm.DB
}

type PageRepositoryUseCase interface {
	GetPage(ctx context.Context) ([]*entity.Page, int64, error)
}

func NewPageRepository(db *gorm.DB) *PageRepository {
	return &PageRepository{db}
}

func (nr *PageRepository) GetPage(ctx context.Context) ([]*entity.Page, int64, error) {
	page := make([]*entity.Page, 0)
	var total int64
	var gormDB = nr.db.
		WithContext(ctx).
		Model(&entity.Page{})

	gormDB = gormDB.Order("created_at ASC")

	if err := gormDB.Find(&page).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[PageRepository-GetPage] error while retrieving page data")
	}

	return page, total, nil
}
