package category_pg

import (
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
