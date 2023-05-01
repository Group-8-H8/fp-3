package user_repository

import (
	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
)

type UserRepository interface {
	Register(payload entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(payload entity.User) (*entity.User, errs.MessageErr)
	UpdateAccount(payload entity.User) (*entity.User, errs.MessageErr)
	DeleteAccount(userId uint) errs.MessageErr
}
