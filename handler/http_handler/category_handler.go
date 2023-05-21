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

// CreateCategory godoc
// @Summary Create Category
// @Description Create a new category
// @Tags category
// @ID create-new-category
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewCreateCategoryRequest true "request body json"
// @Success 201 {object} dto.NewCreateCategoryResponse
// @Router /categories [post]
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

// UpdateCategory godoc
// @Summary Update Category
// @Description Update category name
// @Tags category
// @ID update-category
// @Accept json
// @Produce json
// @Param categoryId path int true "Id of the category"
// @Param RequestBody body dto.NewUpdateCategoryRequest true "request body json"
// @Success 200 {object} dto.NewUpdateCategoryResponse
// @Router /categories/{categoryId} [patch]
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

// GetAllCategories godoc
// @Summary Get All Categories
// @Description Get all categories
// @Tags category
// @ID get-all-category
// @Produce json
// @Success 200 {object} []dto.NewGetCategoriesResponse
// @Router /categories [get]
func (c *categoryHandler) GetCategories(ctx *gin.Context) {
	user := ctx.MustGet("user")

	response, err := c.categoryService.GetCategories(user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetCategory godoc
// @Summary Get Category
// @Description Get category by categories ID
// @Tags category
// @ID get-category
// @Produce json
// @Param categoryId path int true "Id of the category"
// @Success 200 {object} dto.NewGetCategoriesResponse
// @Router /categories/{categoryId} [get]
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

// DeleteCategory godoc
// @Summary Delete Category
// @Description Delete Category
// @Tags category
// @ID delete-category
// @Produce json
// @Param categoryId path int true "Id of the category"
// @Success 200 {object} dto.NewDeleteCategoryResponse
// @Router /categories/{categoryId} [delete]
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
