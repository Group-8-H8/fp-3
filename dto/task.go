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

type NewGetTaskResponse struct {
	Id          int                   `json:"id"`
	Title       string                `json:"title"`
	Status      bool                  `json:"status"`
	Description string                `json:"description"`
	UserId      int                   `json:"user_id"`
	CategoryId  int                   `json:"category_id"`
	CreatedAt   time.Time             `json:"created_at"`
	User        NewUserOnTaskResponse `json:"User"`
}

type NewUpdateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (t *NewUpdateTaskRequest) UpdateTaskRequestToEntity(taskId int, userId int) entity.Task {
	return entity.Task{
		ID:          uint(taskId),
		Title:       t.Title,
		Description: t.Description,
		UserID:      uint(userId),
		UpdatedAt:   time.Now(),
	}
}

type NewUpdateTaskResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"user_id"`
	CategoryId  int       `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
