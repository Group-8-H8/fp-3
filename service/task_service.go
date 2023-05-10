package service

import (
	"github.com/Group-8-H8/fp-3/dto"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/task_repository"
)

type TaskService interface {
	CreateTask(payload dto.NewCreateTaskRequest, userId int) (*dto.NewCreateTaskResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo task_repository.TaskRepository
}

func NewTaskService(taskRepo task_repository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (t *taskService) CreateTask(payload dto.NewCreateTaskRequest, userId int) (*dto.NewCreateTaskResponse, errs.MessageErr) {
	task := payload.CreateTaskRequestToEntity(userId)

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
