package repository

import (
	"context"

	"gorm.io/gorm"

	"gin-starter/entity"
)

// EmailSent struct
type EmailSent struct {
	gormDB *gorm.DB
}

// NewEmailSent will create new email sent repository
func NewEmailSent(db *gorm.DB) *EmailSent {
	return &EmailSent{db}
}

// Insert will insert notification email sent to database
func (r *EmailSent) Insert(ctx context.Context, emailSent *entity.EmailSent) error {
	// check if exist
	exist := &entity.EmailSent{}
	if err := r.gormDB.
		WithContext(ctx).
		Where("m_id = ?", emailSent.MId).
		First(exist).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// insert into database
			if err := r.gormDB.
				WithContext(ctx).
				Model(&entity.EmailSent{}).
				Create(emailSent).
				Error; err != nil {
				return err
			}
		} else {
			// update status
			if err := r.gormDB.
				WithContext(ctx).
				Model(&entity.EmailSent{}).
				Where("m_id = ?", emailSent.MId).
				Update("status", "OUTGOING").
				Error; err != nil {
				return err
			}
		}
	}

	return nil
}

// UpdateStatus will update status of email sent
func (r *EmailSent) UpdateStatus(ctx context.Context, emailSent *entity.EmailSent) error {
	return r.gormDB.
		WithContext(ctx).
		Model(&entity.EmailSent{}).
		Where("m_id = ?", emailSent.MId).
		Update("status", emailSent.Status).
		Error
}
