package resource

import (
	"gin-starter/entity"

	"github.com/google/uuid"
)

type GetFaqResponse struct {
	List  []*Faq `json:"list"`
	Total int64  `json:"total"`
}

type GetFaqByID struct {
	ID string `uri:"id" binding:"required"`
}

type CreateFaqRequest struct {
	Question string `form:"question" json:"question" binding:"required"`
	Answer   string `form:"answer" json:"answer" binding:"required"`
}

type UpdateFaqRequest struct {
	Question string `form:"question" json:"question" binding:"required"`
	Answer   string `form:"answer" json:"answer" binding:"required"`
}

type DeleteFaqRequest struct {
	ID string `uri:"id" binding:"required"`
}

type Faq struct {
	ID       uuid.UUID `json:"id"`
	Question string    `json:"question"`
	Answer   string    `json:"answer"`
}

func NewFaqResponse(faq *entity.Faq) *Faq {
	return &Faq{
		ID:       faq.ID,
		Question: faq.Question,
		Answer:   faq.Answer,
	}
}
