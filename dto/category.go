package dto

import (
	"time"

	"github.com/Group-8-H8/fp-3/entity"
)

type NewCreateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

func (c *NewCreateCategoryRequest) CreateCategoryRequestToEntity() entity.Category {
	return entity.Category{
		Type:      c.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type NewCreateCategoryResponse struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}
