package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rizabach29/todolist-go/app"
)

type IJWTService interface {
	GenerateToken(email string, role string) string
}

type authCustomClaims struct {
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
}

func NewJWTService() IJWTService {
	return &jwtServices{
		secretKey: app.GetEnv("SECRET"),
	}
}

func (service *jwtServices) GenerateToken(email string, role string) string {
	claims := &authCustomClaims{
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}