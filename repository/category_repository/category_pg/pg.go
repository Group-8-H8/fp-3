package category_pg

import (
	"errors"

	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/category_repository"

	"gorm.io/gorm"
)

type categoryPg struct {
	db *gorm.DB
}

func NewCategoryPg(db *gorm.DB) category_repository.CategoryRepository {
	return &categoryPg{db: db}
}

func (c *categoryPg) CreateCategory(payload entity.Category) (*entity.Category, errs.MessageErr) {
	if err := c.db.Create(&payload).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (c *categoryPg) UpdateCategory(payload entity.Category) (*entity.Category, errs.MessageErr) {
	err := c.db.Model(&payload).Where("id = ?", payload.ID).Updates(entity.Category{Type: payload.Type, UpdatedAt: payload.UpdatedAt}).Error
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	return &payload, nil
}

func (c *categoryPg) GetCategories() ([]entity.Category, errs.MessageErr) {
	var categories []entity.Category

	if err := c.db.Preload("Tasks").Find(&categories).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return categories, nil
}

func (c *categoryPg) GetCategory(categoryId int) (*entity.Category, errs.MessageErr) {
	var category entity.Category

	if err := c.db.Preload("Tasks").First(&category, "categories.id = ?", categoryId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("category not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &category, nil
}

func (c *categoryPg) DeleteCategory(categoryId int) errs.MessageErr {
	category := entity.Category{}

	if err := c.db.Where("id = ?", categoryId).Delete(&category).Error; err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
