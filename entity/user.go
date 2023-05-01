package entity

import (
	"os"
	"strings"
	"time"

	"github.com/Group-8-H8/fp-3/pkg/errs"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Full_name string `gorm:"not null"`
	Email     string `gorm:"not null;unique;type:varchar(191)"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
	Tasks     []Task
	CreatedAt time.Time
	UpdatedAt time.Time
}

var secretKey = os.Getenv("SECRET_KEY")

func (u *User) HashPassword() errs.MessageErr {
	const cost = 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), cost)
	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	u.Password = string(bs)
	return nil
}

func (u *User) ComparePassword(reqPassword string) errs.MessageErr {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(reqPassword)); err != nil {
		return errs.NewBadRequestError("wrong password")
	}

	return nil
}

func (u *User) tokenClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	secretKey := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(secretKey))

	return signedToken
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaims()

	return u.signToken(claims)
}

func (u *User) parseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthenticatedError("invalid token error")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, errs.NewUnauthenticatedError("invalid token error")
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claims jwt.MapClaims) errs.MessageErr {
	if id, ok := claims["id"].(float64); !ok {
		return errs.NewUnauthenticatedError("invalid token error")
	} else {
		u.ID = uint(id)
	}

	if email, ok := claims["email"].(string); !ok {
		return errs.NewUnauthenticatedError("invalid token error")
	} else {
		u.Email = email
	}

	return nil
}

func (u *User) VerifyToken(bearerToken string) errs.MessageErr {
	bearer := strings.HasPrefix(bearerToken, "Bearer")
	if !bearer {
		return errs.NewUnauthenticatedError("invalid token error")
	}

	tokenString := strings.Split(bearerToken, " ")[1]

	token, err := u.parseToken(tokenString)
	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errs.NewUnauthenticatedError("invalid token error")
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)
	if err != nil {
		return err
	}

	return nil

}
