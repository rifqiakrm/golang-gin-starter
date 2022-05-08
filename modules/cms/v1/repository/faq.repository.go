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

type FaqRepository struct {
	db *gorm.DB
}

type FaqRepositoryUseCase interface {
	Create(ctx context.Context, faq *entity.Faq) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Faq, error)
	FindAll(ctx context.Context) ([]*entity.Faq, error)
	Update(ctx context.Context, faq *entity.Faq) error
	Delete(ctx context.Context, id uuid.UUID) error
}

func NewFaqRepository(db *gorm.DB) *FaqRepository {
	return &FaqRepository{db}
}

func (nr *FaqRepository) Create(ctx context.Context, faq *entity.Faq) error {
	if err := nr.db.
		WithContext(ctx).
		Model(&entity.Faq{}).
		Create(faq).
		Error; err != nil {
		return errors.Wrap(err, "[FaqRepository-CreateFaq] error while creating user")
	}

	return nil
}

func (nr *FaqRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Faq, error) {
	faq := &entity.Faq{}

	if err := nr.db.
		WithContext(ctx).
		Model(&entity.Faq{}).
		Where("id = ?", id).
		First(&faq).
		Error; err != nil {
		return nil, errors.Wrap(err, "[FaqRepositoryRepository-FindByID] error while getting faq category")
	}

	return faq, nil
}

func (nr *FaqRepository) FindAll(ctx context.Context) ([]*entity.Faq, error) {
	faqs := make([]*entity.Faq, 0)

	if err := nr.db.
		WithContext(ctx).
		Model(&entity.Faq{}).
		Order("created_at asc").
		Find(&faqs).
		Error; err != nil {
		return nil, errors.Wrap(err, "[FaqRepositoryRepository-FindByID] error while getting faq category")
	}

	return faqs, nil
}

func (nr *FaqRepository) Update(ctx context.Context, faq *entity.Faq) error {
	oldTime := faq.UpdatedAt
	faq.UpdatedAt = time.Now()
	if err := nr.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.Faq)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, faq.ID).Error; err != nil {
				log.Println("[GamPTKRepository - Update]", err)
				return err
			}
			if err := tx.Model(&entity.Faq{}).
				Where(`id`, faq.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(faq)).Error; err != nil {
				log.Println("[GamPTKRepository - Update]", err)
				return err
			}
			return nil
		}); err != nil {
		faq.UpdatedAt = oldTime
	}
	return nil
}

func (nr *FaqRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := nr.db.WithContext(ctx).
		Model(&entity.Faq{}).
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
