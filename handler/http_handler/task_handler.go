package http_handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Group-8-H8/fp-3/dto"
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
		errCasting, ok := err.(validator.ValidationErrors)
		if !ok {
			newErrBind := errs.NewBadRequestError("invalid body request")
			ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
			return
		}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBind := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	user := ctx.MustGet("user")

	response, errCreated := t.taskService.CreateTask(requestBody, user)
	if errCreated != nil {
		ctx.AbortWithStatusJSON(errCreated.Status(), errCreated)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (t *taskHandler) GetTasks(ctx *gin.Context) {
	user := ctx.MustGet("user")

	response, err := t.taskService.GetTasks(user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *taskHandler) GetTask(ctx *gin.Context) {
	user := ctx.MustGet("user")

	param := ctx.Param("taskId")
	taskId, errConv := strconv.Atoi(param)
	if errConv != nil {
		newErrConv := errs.NewBadRequestError("invalid task's id")
		ctx.AbortWithStatusJSON(newErrConv.Status(), newErrConv)
		return
	}

	response, err := t.taskService.GetTask(taskId, user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *taskHandler) UpdateTask(ctx *gin.Context) {
	param := ctx.Param("taskId")
	taskId, errParam := strconv.Atoi(param)
	if errParam != nil {
		newErrParam := errs.NewBadRequestError("invalid task's id")
		ctx.AbortWithStatusJSON(newErrParam.Status(), newErrParam)
		return
	}

	var requestBody dto.NewUpdateTaskRequest
	if errBinding := ctx.ShouldBindJSON(&requestBody); errBinding != nil {
		errBinds := []string{}
		errCasting, ok := errBinding.(validator.ValidationErrors)
		if !ok {
			newErrBind := errs.NewBadRequestError("invalid body request")
			ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
			return
		}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("error on field %s, condition : %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBinds := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBinds.Status(), newErrBinds)
		return
	}

	user := ctx.MustGet("user")

	response, err := t.taskService.UpdateTask(requestBody, taskId, user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *taskHandler) UpdateTasksStatus(ctx *gin.Context) {
	param := ctx.Param("taskId")
	taskId, errConv := strconv.Atoi(param)
	if errConv != nil {
		newErrConv := errs.NewBadRequestError("invalid task's id")
		ctx.AbortWithStatusJSON(newErrConv.Status(), newErrConv)
		return
	}

	var requestBody dto.NewUpdateTasksStatusRequest

	if errBinding := ctx.ShouldBindJSON(&requestBody); errBinding != nil {
		errBinds := []string{}
		errCasting, ok := errBinding.(validator.ValidationErrors)
		if !ok {
			newErrBind := errs.NewBadRequestError("invalid body request")
			ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
			return
		}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("error on field %s, condition : %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBinds := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBinds.Status(), newErrBinds)
		return
	}

	user := ctx.MustGet("user")

	response, errResponse := t.taskService.UpdateTasksStatus(requestBody, taskId, user)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *taskHandler) UpdateTasksCategory(ctx *gin.Context) {
	param := ctx.Param("taskId")
	taskId, errConv := strconv.Atoi(param)
	if errConv != nil {
		newError := errs.NewBadRequestError("invalid task's id")
		ctx.AbortWithStatusJSON(newError.Status(), newError)
		return
	}

	var requestBody dto.NewUpdateTasksCategoryRequest

	if errBinding := ctx.ShouldBindJSON(&requestBody); errBinding != nil {
		errBinds := []string{}
		errCasting, ok := errBinding.(validator.ValidationErrors)
		if !ok {
			newErrBind := errs.NewBadRequestError("invalid body request")
			ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
			return
		}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("error on field %s, condition : %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBind := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	user := ctx.MustGet("user")

	response, errResponse := t.taskService.UpdateTasksCategory(requestBody, taskId, user)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (t *taskHandler) DeleteTask(ctx *gin.Context) {
	param := ctx.Param("taskId")
	taskId, errConv := strconv.Atoi(param)
	if errConv != nil {
		newErrConv := errs.NewBadRequestError("invalid task's id")
		ctx.AbortWithStatusJSON(newErrConv.Status(), newErrConv)
		return
	}

	user := ctx.MustGet("user")

	response, errDel := t.taskService.DeleteTask(taskId, user)
	if errDel != nil {
		ctx.AbortWithStatusJSON(errDel.Status(), errDel)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
