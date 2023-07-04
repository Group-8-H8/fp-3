package user_pg

import (
	"fmt"

	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/user_repository"

	"gorm.io/gorm"
)

type userPG struct {
	db *gorm.DB
}

func NewUserPG(db *gorm.DB) user_repository.UserRepository {
	return &userPG{db: db}
}

func (u *userPG) Register(payload entity.User) (*entity.User, errs.MessageErr) {
	if err := u.db.Create(&payload).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &payload, nil
}

func (u *userPG) GetUserByEmail(payload entity.User) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		return nil, errs.NewNotFoundError(fmt.Sprintf("user with email %s is not found", payload.Email))
	}

	return &user, nil
}

func (u *userPG) GetUserById(userId int) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, userId).Error; err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) UpdateAccount(payload entity.User) (*entity.User, errs.MessageErr) {
	var user entity.User

	err := u.db.Model(&user).Where("id = ?", payload.ID).Updates(entity.User{Email: payload.Email, Full_name: payload.Full_name, UpdatedAt: payload.UpdatedAt}).Error
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (u *userPG) DeleteAccount(userId uint) errs.MessageErr {
	user := entity.User{}

	if err := u.db.Where("id = ?", userId).Delete(&user).Error; err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (u *userPG) DeleteAccountByEmail(email string) errs.MessageErr {
	user := entity.User{}

	if err := u.db.Where("email = ?", email).Delete(&user).Error; err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
