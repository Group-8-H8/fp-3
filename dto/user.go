package dto

import (
	"time"

	"github.com/Group-8-H8/fp-3/entity"
)

type NewRegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *NewRegisterRequest) RegisterRequestToEntity() entity.User {
	return entity.User{
		Full_name: u.FullName,
		Email:     u.Email,
		Password:  u.Password,
		Role:      "member",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type NewRegisterResponse struct {
	Id        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type NewLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (u *NewLoginRequest) LoginRequestToEntity() entity.User {
	return entity.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

type NewLoginResponse struct {
	Token string `json:"token"`
}

type NewUpdateAccountRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func (u *NewUpdateAccountRequest) UpdateAccountRequestToEntity(id uint) entity.User {
	return entity.User{
		ID:        id,
		Email:     u.Email,
		Full_name: u.FullName,
		UpdatedAt: time.Now(),
	}
}

type NewUpdateAccountResponse struct {
	Id        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewDeleteAccountResponse struct {
	Message string `json:"message"`
}

type NewUserOnTaskResponse struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
