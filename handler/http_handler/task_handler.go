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

// CreateTask godoc
// @Summary Create Task
// @Description Create a new task
// @Tags task
// @ID create-new-task
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewCreateTaskRequest true "request body json"
// @Success 201 {object} dto.NewCreateTaskResponse
// @Router /tasks [post]
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

// GetAllTasks godoc
// @Summary Get All Tasks
// @Description Get all tasks
// @Tags task
// @ID get-all-task
// @Produce json
// @Success 200 {object} []dto.NewGetTaskResponse
// @Router /tasks [get]
func (t *taskHandler) GetTasks(ctx *gin.Context) {
	user := ctx.MustGet("user")

	response, err := t.taskService.GetTasks(user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetTasksById godoc
// @Summary Get Task By ID
// @Description Get task by task's ID
// @Tags task
// @ID get-task-by-id
// @Produce json
// @Param taskId path int true "Id of the task"
// @Success 200 {object} dto.NewGetTaskResponse
// @Router /tasks/{taskId} [get]
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

// UpdateTask godoc
// @Summary Update Task
// @Description Update task's detail
// @Tags task
// @ID update-task
// @Accept json
// @Produce json
// @Param taskId path int true "Id of the task"
// @Param requestBody body dto.NewUpdateTaskRequest true "request body json"
// @Success 200 {object} dto.NewUpdateTaskResponse
// @Router /tasks/{taskId} [put]
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

// UpdateTasksStatus godoc
// @Summary Update Task's Status
// @Description Update task's status
// @Tags task
// @ID update-tasks-status
// @Accept json
// @Produce json
// @Param taskId path int true "Id of the task"
// @Param requestBody body dto.NewUpdateTasksStatusRequest true "request body json"
// @Success 200 {object} dto.NewUpdateTaskResponse
// @Router /tasks/update-status/{taskId} [patch]
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

// UpdateTasksCategory godoc
// @Summary Update Task's Category
// @Description Update task's category
// @Tags task
// @ID update-tasks-category
// @Accept json
// @Produce json
// @Param taskId path int true "Id of the task"
// @Param requestBody body dto.NewUpdateTasksCategoryRequest true "request body json"
// @Success 200 {object} dto.NewUpdateTaskResponse
// @Router /tasks/update-category/{taskId} [patch]
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

// DeleteTask godoc
// @Summary Delete Task
// @Description Delete task by task's ID
// @Tags task
// @ID delete-task
// @Produce json
// @Param taskId path int true "Id of the task"
// @Success 200 {object} dto.NewDeleteTaskResponse
// @Router /tasks/{taskId} [delete]
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
