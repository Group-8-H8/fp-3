package task_pg

import (
	"errors"

	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/task_repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) task_repository.TaskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) CreateTask(payload entity.Task) (*entity.Task, errs.MessageErr) {
	if err := t.db.Create(&payload).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (t *taskRepository) GetTasks(userId int) ([]entity.Task, errs.MessageErr) {
	var tasks []entity.Task

	if err := t.db.Where("user_id = ?", userId).Find(&tasks).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return tasks, nil
}

func (t *taskRepository) GetTask(taskId int, userId int) (*entity.Task, errs.MessageErr) {
	var task entity.Task

	if err := t.db.First(&task, "id = ? AND user_id = ?", taskId, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("task not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &task, nil
}

func (t *taskRepository) UpdateTask(payload entity.Task) (*entity.Task, errs.MessageErr) {
	err := t.db.Model(&payload).Where("id = ? AND user_id = ?", payload.ID, payload.UserID).Updates(entity.Task{Title: payload.Title, Description: payload.Description, UpdatedAt: payload.UpdatedAt}).Error

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (t *taskRepository) UpdateTasksStatus(payload entity.Task) (*entity.Task, errs.MessageErr) {
	err := t.db.Model(&payload).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", payload.ID, payload.UserID).Update("status", payload.Status).Error

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (t *taskRepository) UpdateTasksCategory(payload entity.Task) (*entity.Task, errs.MessageErr) {
	err := t.db.Model(&payload).Clauses(clause.Returning{}).Where("id = ? AND user_id = ?", payload.ID, payload.UserID).Update("category_id", payload.CategoryID).Error

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (t *taskRepository) DeleteTask(taskId int, userId int) errs.MessageErr {
	task := entity.Task{}

	if err := t.db.Where("id = ? AND user_id = ?", taskId, userId).Delete(&task).Error; err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
