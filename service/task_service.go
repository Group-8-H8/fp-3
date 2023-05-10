package service

import (
	"github.com/Group-8-H8/fp-3/dto"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/category_repository"
	"github.com/Group-8-H8/fp-3/repository/task_repository"
)

type TaskService interface {
	CreateTask(payload dto.NewCreateTaskRequest, userId int) (*dto.NewCreateTaskResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo     task_repository.TaskRepository
	categoryRepo category_repository.CategoryRepository
}

func NewTaskService(taskRepo task_repository.TaskRepository, categoryRepo category_repository.CategoryRepository) TaskService {
	return &taskService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
	}
}

func (t *taskService) CreateTask(payload dto.NewCreateTaskRequest, userId int) (*dto.NewCreateTaskResponse, errs.MessageErr) {
	task := payload.CreateTaskRequestToEntity(userId)

	if _, err := t.categoryRepo.GetCategory(int(task.CategoryID)); err != nil && err.Status() == 404 {
		return nil, errs.NewNotFoundError("invalid category")
	}

	createdTask, err := t.taskRepo.CreateTask(task)
	if err != nil {
		return nil, err
	}

	response := &dto.NewCreateTaskResponse{
		Id:          int(createdTask.ID),
		Title:       createdTask.Title,
		Status:      createdTask.Status,
		Description: createdTask.Description,
		UserId:      int(createdTask.UserID),
		CategoryId:  int(createdTask.CategoryID),
		CreatedAt:   createdTask.CreatedAt,
	}

	return response, nil
}
