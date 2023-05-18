package service

import (
	"time"

	"github.com/Group-8-H8/fp-3/dto"
	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/category_repository"
	"github.com/Group-8-H8/fp-3/repository/task_repository"
	"github.com/Group-8-H8/fp-3/repository/user_repository"
)

type TaskService interface {
	CreateTask(payload dto.NewCreateTaskRequest, payloadUser any) (*dto.NewCreateTaskResponse, errs.MessageErr)
	GetTasks(payload any) ([]dto.NewGetTaskResponse, errs.MessageErr)
	GetTask(taskId int, payloadUser any) (*dto.NewGetTaskResponse, errs.MessageErr)
	UpdateTask(payload dto.NewUpdateTaskRequest, taskId int, payloadUser any) (*dto.NewUpdateTaskResponse, errs.MessageErr)
	UpdateTasksStatus(payload dto.NewUpdateTasksStatusRequest, taskId int, payloadUser any) (*dto.NewUpdateTaskResponse, errs.MessageErr)
	UpdateTasksCategory(payload dto.NewUpdateTasksCategoryRequest, taskId int, payloadUser any) (*dto.NewUpdateTaskResponse, errs.MessageErr)
	DeleteTask(taskId int, payloadUser any) (*dto.NewDeleteTaskResponse, errs.MessageErr)
}

type taskService struct {
	taskRepo     task_repository.TaskRepository
	categoryRepo category_repository.CategoryRepository
	userRepo     user_repository.UserRepository
}

func NewTaskService(taskRepo task_repository.TaskRepository, categoryRepo category_repository.CategoryRepository, userRepo user_repository.UserRepository) TaskService {
	return &taskService{
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
	}
}

func (t *taskService) CreateTask(payload dto.NewCreateTaskRequest, payloadUser any) (*dto.NewCreateTaskResponse, errs.MessageErr) {
	user := payloadUser.(entity.User)

	task := payload.CreateTaskRequestToEntity(int(user.ID))

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

func (t *taskService) GetTasks(payload any) ([]dto.NewGetTaskResponse, errs.MessageErr) {
	userPayload := payload.(entity.User)

	getTasks, err := t.taskRepo.GetTasks(int(userPayload.ID))
	if err != nil {
		return nil, err
	}

	user, _ := t.userRepo.GetUserById(int(userPayload.ID))
	userResponse := dto.NewUserOnTaskResponse{
		Id:       int(user.ID),
		Email:    user.Email,
		FullName: user.Full_name,
	}

	responses := []dto.NewGetTaskResponse{}
	for _, task := range getTasks {
		response := dto.NewGetTaskResponse{
			Id:          int(task.ID),
			Title:       task.Title,
			Status:      task.Status,
			Description: task.Description,
			UserId:      int(task.UserID),
			CategoryId:  int(task.CategoryID),
			CreatedAt:   task.CreatedAt,
			User:        userResponse,
		}

		responses = append(responses, response)
	}

	return responses, nil
}

func (t *taskService) GetTask(taskId int, payloadUser any) (*dto.NewGetTaskResponse, errs.MessageErr) {
	userPayload := payloadUser.(entity.User)
	user, _ := t.userRepo.GetUserById(int(userPayload.ID))
	userResponse := dto.NewUserOnTaskResponse{
		Id:       int(user.ID),
		Email:    user.Email,
		FullName: user.Full_name,
	}

	getTask, err := t.taskRepo.GetTask(taskId, int(userPayload.ID))
	if err != nil {
		return nil, err
	}

	response := &dto.NewGetTaskResponse{
		Id:          int(getTask.ID),
		Title:       getTask.Title,
		Status:      getTask.Status,
		Description: getTask.Description,
		UserId:      int(getTask.UserID),
		CategoryId:  int(getTask.CategoryID),
		CreatedAt:   getTask.CreatedAt,
		User:        userResponse,
	}

	return response, nil
}

func (t *taskService) UpdateTask(payload dto.NewUpdateTaskRequest, taskId int, payloadUser any) (*dto.NewUpdateTaskResponse, errs.MessageErr) {
	user := payloadUser.(entity.User)

	task := payload.UpdateTaskRequestToEntity(taskId, int(user.ID))

	if _, errCheck := t.taskRepo.GetTask(taskId, int(user.ID)); errCheck != nil && errCheck.Status() == 404 {
		return nil, errCheck
	}

	updatedTask, err := t.taskRepo.UpdateTask(task)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUpdateTaskResponse{
		Id:          int(updatedTask.ID),
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserId:      int(updatedTask.UserID),
		CategoryId:  int(updatedTask.CategoryID),
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) UpdateTasksStatus(payload dto.NewUpdateTasksStatusRequest, taskId int, payloadUser any) (*dto.NewUpdateTaskResponse, errs.MessageErr) {
	user := payloadUser.(entity.User)

	task := entity.Task{
		ID:        uint(taskId),
		Status:    payload.Status,
		UserID:    uint(user.ID),
		UpdatedAt: time.Now(),
	}

	if _, errCheck := t.taskRepo.GetTask(taskId, int(user.ID)); errCheck != nil && errCheck.Status() == 404 {
		return nil, errCheck
	}

	updatedTask, err := t.taskRepo.UpdateTasksStatus(task)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUpdateTaskResponse{
		Id:          int(updatedTask.ID),
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserId:      int(updatedTask.UserID),
		CategoryId:  int(updatedTask.CategoryID),
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) UpdateTasksCategory(payload dto.NewUpdateTasksCategoryRequest, taskId int, payloadUser any) (*dto.NewUpdateTaskResponse, errs.MessageErr) {
	user := payloadUser.(entity.User)

	if _, errCheck := t.taskRepo.GetTask(taskId, int(user.ID)); errCheck != nil && errCheck.Status() == 404 {
		return nil, errCheck
	}

	if _, errCat := t.categoryRepo.GetCategory(payload.CategoryId); errCat != nil && errCat.Status() == 404 {
		return nil, errCat
	}

	task := entity.Task{
		ID:         uint(taskId),
		CategoryID: uint(payload.CategoryId),
		UserID:     user.ID,
		UpdatedAt:  time.Now(),
	}

	updatedTask, errUpdate := t.taskRepo.UpdateTasksCategory(task)
	if errUpdate != nil {
		return nil, errUpdate
	}

	response := &dto.NewUpdateTaskResponse{
		Id:          int(updatedTask.ID),
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		UserId:      int(updatedTask.UserID),
		CategoryId:  int(updatedTask.CategoryID),
		UpdatedAt:   updatedTask.UpdatedAt,
	}

	return response, nil
}

func (t *taskService) DeleteTask(taskId int, payloadUser any) (*dto.NewDeleteTaskResponse, errs.MessageErr) {
	user := payloadUser.(entity.User)

	if _, errCheck := t.taskRepo.GetTask(taskId, int(user.ID)); errCheck != nil && errCheck.Status() == 404 {
		return nil, errCheck
	}

	if errDel := t.taskRepo.DeleteTask(taskId, int(user.ID)); errDel != nil {
		return nil, errDel
	}

	response := &dto.NewDeleteTaskResponse{
		Message: "Task has been sucessfully deleted",
	}

	return response, nil
}
