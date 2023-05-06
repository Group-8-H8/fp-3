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

type NewUpdateCategoryRequest struct {
	Type string `json:"type" binding:"required"`
}

func (c *NewUpdateCategoryRequest) UpdateCategoryRequestToEntity(categoryId uint) entity.Category {
	return entity.Category{
		ID:        categoryId,
		Type:      c.Type,
		UpdatedAt: time.Now(),
	}
}

type NewUpdateCategoryResponse struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updated_at"`
}
