package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	pageTableName = "pages"
)

type Page struct {
	ID      uuid.UUID `gorm:"primaryKey;" json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Auditable
}

func NewPage(
	id uuid.UUID,
	title string,
	content string,
	createdBy string,
) *Page {
	return &Page{
		ID:        id,
		Title:     title,
		Content:   content,
		Auditable: NewAuditable(createdBy),
	}
}

// TableName specifies table name
func (model *Page) TableName() string {
	return pageTableName
}

// MapUpdateFrom mapping from model
func (model *Page) MapUpdateFrom(from *Page) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"title":      model.Title,
			"content":    model.Content,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Title != from.Title {
		mapped["title"] = from.Title
	}

	if model.Content != from.Content {
		mapped["content"] = from.Content
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
