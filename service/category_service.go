package service

import (
	"github.com/Group-8-H8/fp-3/dto"
	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/category_repository"
)

type CategoryService interface {
	CreateCategory(payload dto.NewCreateCategoryRequest) (*dto.NewCreateCategoryResponse, errs.MessageErr)
	UpdateCategory(payload dto.NewUpdateCategoryRequest, categoryId uint) (*dto.NewUpdateCategoryResponse, errs.MessageErr)
	GetCategories(payloadUser any) ([]dto.NewGetCategoriesResponse, errs.MessageErr)
	GetCategory(categoryId int, payloadUser any) (*dto.NewGetCategoriesResponse, errs.MessageErr)
	DeleteCategory(categoryId int) (*dto.NewDeleteCategoryResponse, errs.MessageErr)
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

	if _, err := c.categoryRepo.GetCategory(int(categoryId)); err != nil && err.Status() == 404 {
		return nil, err
	}

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

func (c *categoryService) GetCategories(payloadUser any) ([]dto.NewGetCategoriesResponse, errs.MessageErr) {
	user := payloadUser.(entity.User)

	getCategories, err := c.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}

	var responses []dto.NewGetCategoriesResponse
	for _, e := range getCategories {
		tasks := []entity.Task{}
		for _, task := range e.Tasks {
			if task.UserID == user.ID {
				tasks = append(tasks, task)
			}
		}
		response := dto.NewGetCategoriesResponse{
			Id:        e.ID,
			Type:      e.Type,
			UpdatedAt: e.UpdatedAt,
			CreatedAt: e.CreatedAt,
			Tasks:     tasks,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (c *categoryService) GetCategory(categoryId int, payloadUser any) (*dto.NewGetCategoriesResponse, errs.MessageErr) {
	user := payloadUser.(entity.User)

	getCategory, err := c.categoryRepo.GetCategory(categoryId)
	if err != nil {
		return nil, err
	}

	tasks := []entity.Task{}
	for _, task := range getCategory.Tasks {
		if task.UserID == user.ID {
			tasks = append(tasks, task)
		}
	}

	response := &dto.NewGetCategoriesResponse{
		Id:        getCategory.ID,
		Type:      getCategory.Type,
		UpdatedAt: getCategory.UpdatedAt,
		CreatedAt: getCategory.CreatedAt,
		Tasks:     tasks,
	}

	return response, nil
}

func (c *categoryService) DeleteCategory(categoryId int) (*dto.NewDeleteCategoryResponse, errs.MessageErr) {
	if _, err := c.categoryRepo.GetCategory(categoryId); err != nil && err.Status() == 404 {
		return nil, err
	}

	if err := c.categoryRepo.DeleteCategory(categoryId); err != nil {
		return nil, err
	}

	response := &dto.NewDeleteCategoryResponse{
		Message: "Category has been successfully deleted",
	}

	return response, nil
}
