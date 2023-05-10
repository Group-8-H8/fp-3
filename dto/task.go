package dto

import (
	"time"

	"github.com/Group-8-H8/fp-3/entity"
)

type NewCreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryId  int    `json:"category_id" binding:"number"`
}

func (t *NewCreateTaskRequest) CreateTaskRequestToEntity(userId int) entity.Task {
	return entity.Task{
		Title:       t.Title,
		Description: t.Description,
		Status:      false,
		CategoryID:  uint(t.CategoryId),
		UserID:      uint(userId),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

type NewCreateTaskResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
