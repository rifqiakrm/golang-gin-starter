package repository

import (
	"context"
	"gin-starter/entity"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PageRepository struct {
	db *gorm.DB
}

type PageRepositoryUseCase interface {
	Create(ctx context.Context, page *entity.Page) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Page, error)
	FindAll(ctx context.Context) ([]*entity.Page, error)
	Update(ctx context.Context, page *entity.Page) error
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewPageRepository(db *gorm.DB) *PageRepository {
	return &PageRepository{db}
}

func (nr *PageRepository) Create(ctx context.Context, page *entity.Page) error {
	if err := nr.db.
		WithContext(ctx).
		Model(&entity.Page{}).
		Create(page).
		Error; err != nil {
		return errors.Wrap(err, "[PageRepository-CreatePage] error while creating user")
	}

	return nil
}

func (nr *PageRepository) FindAll(ctx context.Context) ([]*entity.Page, error) {
	page := make([]*entity.Page, 0)

	if err := nr.db.
		WithContext(ctx).
		Model(&entity.Page{}).
		Find(&page).
		Error; err != nil {
		return nil, errors.Wrap(err, "[PageRepositoryRepository-FindByID] error while getting page category")
	}

	return page, nil
}

func (nr *PageRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Page, error) {
	page := &entity.Page{}

	if err := nr.db.
		WithContext(ctx).
		Model(&entity.Page{}).
		Where("id = ?", id).
		First(&page).
		Error; err != nil {
		return nil, errors.Wrap(err, "[PageRepositoryRepository-FindByID] error while getting page category")
	}

	return page, nil
}

func (nr *PageRepository) Update(ctx context.Context, page *entity.Page) error {
	oldTime := page.UpdatedAt
	page.UpdatedAt = time.Now()
	if err := nr.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.Page)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, page.ID).Error; err != nil {
				log.Println("[GamPTKRepository - Update]", err)
				return err
			}
			if err := tx.Model(&entity.Page{}).
				Where(`id`, page.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(page)).Error; err != nil {
				log.Println("[GamPTKRepository - Update]", err)
				return err
			}
			return nil
		}); err != nil {
		page.UpdatedAt = oldTime
	}
	return nil
}

func (nr *PageRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := nr.db.WithContext(ctx).
		Model(&entity.Page{}).
		Where(`id = ?`, id).
		Updates(
			map[string]interface{}{
				"updated_at": time.Now(),
				"deleted_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeactivateUser] error when updating user data")
	}

	return nil
}
