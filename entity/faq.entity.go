package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	faqTableName = "faqs"
)

type Faq struct {
	ID       uuid.UUID `gorm:"primaryKey;" json:"id"`
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
	Auditable
}

func NewFaq(
	id uuid.UUID,
	question string,
	answer string,
	createdBy string,
) *Faq {
	return &Faq{
		ID:        id,
		Question:  question,
		Answer:    answer,
		Auditable: NewAuditable(createdBy),
	}
}

// TableName specifies table name
func (model *Faq) TableName() string {
	return faqTableName
}

// MapUpdateFrom mapping from model
func (model *Faq) MapUpdateFrom(from *Faq) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"question":   model.Question,
			"answer":     model.Answer,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.Question != from.Question {
		mapped["question"] = from.Question
	}

	if model.Answer != from.Answer {
		mapped["answer"] = from.Answer
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
