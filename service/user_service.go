package service

import (
	"github.com/Group-8-H8/fp-3/dto"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/user_repository"
)

type UserService interface {
	Register(payload dto.NewRegisterRequest) (*dto.NewRegisterResponse, errs.MessageErr)
	Login(payload dto.NewLoginRequest) (*dto.NewLoginResponse, errs.MessageErr)
	UpdateAccount(payload dto.NewUpdateAccountRequest, id uint) (*dto.NewUpdateAccountResponse, errs.MessageErr)
	DeleteAccount(userId uint) (*dto.NewDeleteAccountResponse, errs.MessageErr)
}

type userService struct {
	userRepo user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) Register(payload dto.NewRegisterRequest) (*dto.NewRegisterResponse, errs.MessageErr) {
	user := payload.RegisterRequestToEntity()

	err := user.HashPassword()
	if err != nil {
		return nil, err
	}

	createdUser, err := u.userRepo.Register(user)
	if err != nil {
		return nil, err
	}

	response := &dto.NewRegisterResponse{
		Id:        createdUser.ID,
		FullName:  createdUser.Full_name,
		Email:     createdUser.Email,
		CreatedAt: createdUser.CreatedAt,
	}

	return response, nil
}

func (u *userService) Login(payload dto.NewLoginRequest) (*dto.NewLoginResponse, errs.MessageErr) {
	user := payload.LoginRequestToEntity()

	getUser, err := u.userRepo.GetUserByEmail(user)
	if err != nil {
		return nil, err
	}

	if err := getUser.ComparePassword(user.Password); err != nil {
		return nil, err
	}

	token := getUser.GenerateToken()

	response := &dto.NewLoginResponse{
		Token: token,
	}

	return response, nil
}

func (u *userService) UpdateAccount(payload dto.NewUpdateAccountRequest, id uint) (*dto.NewUpdateAccountResponse, errs.MessageErr) {
	user := payload.UpdateAccountRequestToEntity(id)

	updateUser, err := u.userRepo.UpdateAccount(user)
	if err != nil {
		return nil, err
	}

	response := &dto.NewUpdateAccountResponse{
		Id:        updateUser.ID,
		FullName:  updateUser.Full_name,
		Email:     updateUser.Email,
		UpdatedAt: updateUser.UpdatedAt,
	}

	return response, nil
}

func (u *userService) DeleteAccount(userId uint) (*dto.NewDeleteAccountResponse, errs.MessageErr) {
	if err := u.userRepo.DeleteAccount(userId); err != nil {
		return nil, err
	}

	response := &dto.NewDeleteAccountResponse{
		Message: "Your account has been successfully deleted",
	}

	return response, nil
}
