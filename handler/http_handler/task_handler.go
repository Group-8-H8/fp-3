package http_handler

import (
	"fmt"
	"net/http"

	"github.com/Group-8-H8/fp-3/dto"
	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type taskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) taskHandler {
	return taskHandler{taskService: taskService}
}

func (t *taskHandler) CreateTask(ctx *gin.Context) {
	var requestBody dto.NewCreateTaskRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBinds := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errBind := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBind := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	user := ctx.MustGet("user").(entity.User)

	response, errCreated := t.taskService.CreateTask(requestBody, int(user.ID))
	if errCreated != nil {
		ctx.AbortWithStatusJSON(errCreated.Status(), errCreated)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
