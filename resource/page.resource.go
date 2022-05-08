package resource

import (
	"gin-starter/entity"

	"github.com/google/uuid"
)

type GetPageResponse struct {
	List  []*Page `json:"list"`
	Total int64   `json:"total"`
}

type GetPageByIDRequest struct {
	ID string `uri:"id" binding:"required"`
}

type Page struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

type CreatePageRequest struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type UpdatePageRequest struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type DeletePageRequest struct {
	ID string `uri:"id" binding:"required"`
}

func NewPageResponse(page *entity.Page) *Page {
	return &Page{
		ID:      page.ID,
		Title:   page.Title,
		Content: page.Content,
	}
}
