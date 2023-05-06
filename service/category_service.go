package service

import (
	"github.com/Group-8-H8/fp-3/dto"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/category_repository"
)

type CategoryService interface {
	CreateCategory(payload dto.NewCreateCategoryRequest) (*dto.NewCreateCategoryResponse, errs.MessageErr)
	UpdateCategory(payload dto.NewUpdateCategoryRequest, categoryId uint) (*dto.NewUpdateCategoryResponse, errs.MessageErr)
}

type categoryService struct {
	categoryRepo category_repository.CategoryRepository
}

func NewCategoryService(categoryRepo category_repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (c *categoryService) CreateCategory(payload dto.NewCreateCategoryRequest) (*dto.NewCreateCategoryResponse, errs.MessageErr) {
	category := payload.CreateCategoryRequestToEntity()

	createdCategory, err := c.categoryRepo.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	response := &dto.NewCreateCategoryResponse{
		Id:        createdCategory.ID,
		Type:      category.Type,
		CreatedAt: category.CreatedAt,
	}

	return response, nil
}

func (c *categoryService) UpdateCategory(payload dto.NewUpdateCategoryRequest, categoryId uint) (*dto.NewUpdateCategoryResponse, errs.MessageErr) {
	category := payload.UpdateCategoryRequestToEntity(categoryId)

	updatedCategory, err := c.categoryRepo.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUpdateCategoryResponse{
		Id:        updatedCategory.ID,
		Type:      updatedCategory.Type,
		UpdatedAt: updatedCategory.UpdatedAt,
	}

	return response, nil
}
