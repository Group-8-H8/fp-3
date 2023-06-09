package task_repository

import (
	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
)

type TaskRepository interface {
	CreateTask(payload entity.Task) (*entity.Task, errs.MessageErr)
	GetTasks(userId int) ([]entity.Task, errs.MessageErr)
	GetTask(taskId int, userId int) (*entity.Task, errs.MessageErr)
	UpdateTask(payload entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTasksStatus(payload entity.Task) (*entity.Task, errs.MessageErr)
	UpdateTasksCategory(payload entity.Task) (*entity.Task, errs.MessageErr)
	DeleteTask(taskId int, userId int) errs.MessageErr
}
