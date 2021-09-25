package services

import (
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/repositories"
)

type IAuthService interface {
	Login(loginModel models.LoginModel) (models.User, error)
	Register(registerModel models.RegisterModel) bool
}

type AuthService struct {
	Repo repositories.Repository
}

// Create new struct
func NewAuthService(Repo repositories.Repository) IAuthService {
	return &AuthService{Repo}
}

// Method
func (s *AuthService) Login(loginModel models.LoginModel) (models.User, error) {
	s.Repo.UserRepostory.GetById(loginModel.Email)
	return models.User{}, nil
}

func (s *AuthService) Register(registerModel models.RegisterModel) bool {
	// if _, err := s.Repo.UserRepostory.GetById(registerModel.Email); err != nil {
	// 	return false
	// }

	registerModel.Password, _ = helpers.HashPassword(registerModel.Password)
	registerModel.ForgotPassword, _ = helpers.HashPassword(registerModel.ForgotPassword)

	// s.Repo.UserRepostory.Create(registerModel)

	return true
}