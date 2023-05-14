package task_pg

import (
	"errors"

	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/task_repository"
	"gorm.io/gorm"
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
