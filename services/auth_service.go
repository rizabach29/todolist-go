package services

import (
	"errors"

	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/middleware"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/repositories"
)

type IAuthService interface {
	Login(loginModel models.LoginModel) (*string, error)
	Register(registerModel models.RegisterModel) error
}

type AuthService struct {
	Repo repositories.Repository
}

// Create new struct
func NewAuthService(Repo repositories.Repository) IAuthService {
	return &AuthService{Repo}
}

// Method
func (s *AuthService) Login(loginModel models.LoginModel) (*string, error) {
	user, err := s.Repo.UserRepostory.GetById(loginModel.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if !helpers.CheckPasswordHash(loginModel.Password, user.Password) {
		return nil, errors.New("password invalid")
	}

	token := middleware.NewJWTService().GenerateToken(user.Email)
	return &token, nil
}

func (s *AuthService) Register(registerModel models.RegisterModel) error {
	if _, err := s.Repo.UserRepostory.GetById(registerModel.Email); err == nil {
		return errors.New("user exist")
	}

	registerModel.Password, _ = helpers.HashPassword(registerModel.Password)
	registerModel.ForgotPassword, _ = helpers.HashPassword(registerModel.ForgotPassword)
	
	s.Repo.UserRepostory.Create(registerModel)
	return nil	
}