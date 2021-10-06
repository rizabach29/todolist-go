package services

import (
	"errors"

	"github.com/rizabach29/todolist-go/helpers"
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
	user, err := s.Repo.UserRepostory.GetByEmail(loginModel.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !helpers.CheckPasswordHash(loginModel.Password, user.Password) {
		return nil, errors.New("password invalid")
	}

	role, _ := s.Repo.RoleRepository.GetById(user.RoleId)

	token := NewJWTService().GenerateToken(user.Email, role.Name)
	return &token, nil
}

func (s *AuthService) Register(registerModel models.RegisterModel) error {
	if _, err := s.Repo.UserRepostory.GetByEmail(registerModel.Email); err == nil {
		return errors.New("user exist")
	}

	registerModel.Password, _ = helpers.HashPassword(registerModel.Password)
	registerModel.ForgotPassword, _ = helpers.HashPassword(registerModel.ForgotPassword)
	
	s.Repo.UserRepostory.Create(registerModel)
	return nil	
}