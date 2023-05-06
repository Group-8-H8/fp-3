package category_repository

import (
	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
)

type CategoryRepository interface {
	CreateCategory(payload entity.Category) (*entity.Category, errs.MessageErr)
	UpdateCategory(payload entity.Category) (*entity.Category, errs.MessageErr)
}
