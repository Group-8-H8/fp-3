package service

import (
	"github.com/Group-8-H8/fp-3/entity"
	"github.com/Group-8-H8/fp-3/pkg/errs"
	"github.com/Group-8-H8/fp-3/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
}

type authService struct {
	userRepo user_repository.UserRepository
}

func NewAuthService(userRepo user_repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if _, err = a.userRepo.GetUserByEmail(user); err != nil {
			errToken := errs.NewUnauthenticatedError("invalid token error")
			ctx.AbortWithStatusJSON(errToken.Status(), errToken)
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

func (a *authService) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(entity.User)

		getUser, err := a.userRepo.GetUserByEmail(user)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if err = errs.NewUnautorizhedError("forbidden"); getUser.Role != "admin" {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		ctx.Next()
	}
}
