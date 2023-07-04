package http_handler

import (
	"fmt"
	"log"
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

// RegisterNewAccount godoc
// @Summary Create Account
// @Description Create a new account
// @Tags user
// @ID register-new-account
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewRegisterRequest true "request body json"
// @Success 201 {object} dto.NewRegisterResponse
// @Router /users/register [post]
func (u *userHandler) Register(ctx *gin.Context) {
	var requestBody dto.NewRegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBinds := []string{}
		errCasting, ok := err.(validator.ValidationErrors)
		if !ok {
			newErrBind := errs.NewBadRequestError("invalid body request")
			ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
			return
		}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}
		newErrBind := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	createdUser, errResponse := u.userService.Register(requestBody)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

// LoginAccount godoc
// @Summary Login Registered Account
// @Description Login registered account to get the token
// @Tags user
// @ID login-account
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewLoginRequest true "request body json"
// @Success 200 {object} dto.NewLoginResponse
// @Router /users/login [post]
func (u *userHandler) Login(ctx *gin.Context) {
	var requestBody dto.NewLoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		errBinds := []string{}
		errCasting, ok := err.(validator.ValidationErrors)
		if !ok {
			newErrBind := errs.NewBadRequestError("invalid body request")
			ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
			return
		}
		for _, e := range errCasting {
			errBind := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errBinds = append(errBinds, errBind)
		}

		newErrBind := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBind.Status(), newErrBind)
		return
	}

	token, errResponse := u.userService.Login(requestBody)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

// UpdateAccount godoc
// @Summary Update Account
// @Description Update account's fullname and email
// @Tags user
// @ID update-account
// @Accept json
// @Produce json
// @Param requestBody body dto.NewUpdateAccountRequest true "request body json"
// @Success 200 {object} dto.NewUpdateAccountResponse
// @Router /users/update-account [put]
func (u *userHandler) UpdateAccount(ctx *gin.Context) {
	user := ctx.MustGet("user").(entity.User)

	var bodyRequest dto.NewUpdateAccountRequest

	if err := ctx.ShouldBindJSON(&bodyRequest); err != nil {
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
	}

	response, errResponse := u.userService.UpdateAccount(bodyRequest, user.ID)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// DeleteAccount godoc
// @Summary Delete Account
// @Description Delete an account
// @Tags user
// @ID delete-account
// @Produce json
// @Success 200 {object} dto.NewDeleteAccountResponse
// @Router /users/delete-account [delete]
func (u *userHandler) DeleteAccount(ctx *gin.Context) {
	user := ctx.MustGet("user").(entity.User)

	response, err := u.userService.DeleteAccount(user.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (u *userHandler) SeedAdminAccount() {
	err := u.userService.SeedAdminAccount()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
