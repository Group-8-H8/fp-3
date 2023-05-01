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

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (u *userHandler) Register(ctx *gin.Context) {
	var requestBody dto.NewRegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBinds := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errBind := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBind := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	createdUser, err := u.userService.Register(requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (u *userHandler) Login(ctx *gin.Context) {
	var requestBody dto.NewLoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBinds := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errBind := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}

		newErrBind := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	token, err := u.userService.Login(requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (u *userHandler) UpdateAccount(ctx *gin.Context) {
	user := ctx.MustGet("user").(entity.User)

	var bodyRequest dto.NewUpdateAccountRequest

	if err := ctx.ShouldBindJSON(&bodyRequest); err != nil {
		errBinds := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errBind := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBind := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
	}

	response, err := u.userService.UpdateAccount(bodyRequest, user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (u *userHandler) DeleteAccount(ctx *gin.Context) {
	user := ctx.MustGet("user").(entity.User)

	response, err := u.userService.DeleteAccount(user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
