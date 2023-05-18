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

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) categoryHandler {
	return categoryHandler{categoryService: categoryService}
}

func (c *categoryHandler) CreateCategory(ctx *gin.Context) {
	var requestBody dto.NewCreateCategoryRequest

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

	response, errResponse := c.categoryService.CreateCategory(requestBody)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *categoryHandler) UpdateCategory(ctx *gin.Context) {
	var requestBody dto.NewUpdateCategoryRequest

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
		newErrBinds := errs.NewUnprocessableEntityError(errBinds)
		ctx.AbortWithStatusJSON(newErrBinds.Status(), newErrBinds)
		return
	}

	param := ctx.Param("categoryId")
	u64, errConv := strconv.ParseUint(param, 10, 32)
	if errConv != nil {
		newErrBadReq := errs.NewBadRequestError("invalid category id")
		ctx.AbortWithStatusJSON(newErrBadReq.Status(), newErrBadReq)
		return
	}
	categoryId := uint(u64)

	response, errResponse := c.categoryService.UpdateCategory(requestBody, categoryId)
	if errResponse != nil {
		ctx.AbortWithStatusJSON(errResponse.Status(), errResponse)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *categoryHandler) GetCategories(ctx *gin.Context) {
	user := ctx.MustGet("user")

	response, err := c.categoryService.GetCategories(user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *categoryHandler) GetCategory(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")

	id, err := strconv.Atoi(categoryId)
	if err != nil {
		errBadReq := errs.NewBadRequestError("invalid category id")
		ctx.AbortWithStatusJSON(errBadReq.Status(), errBadReq)
		return
	}

	user := ctx.MustGet("user")

	response, errGet := c.categoryService.GetCategory(id, user)
	if errGet != nil {
		ctx.AbortWithStatusJSON(errGet.Status(), errGet)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *categoryHandler) DeleteCategory(ctx *gin.Context) {
	param := ctx.Param("categoryId")
	categoryId, errConv := strconv.Atoi(param)
	if errConv != nil {
		newErrConv := errs.NewBadRequestError("invalid category id")
		ctx.AbortWithStatusJSON(newErrConv.Status(), newErrConv)
		return
	}

	response, err := c.categoryService.DeleteCategory(categoryId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
