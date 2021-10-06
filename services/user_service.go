package services

import (
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/repositories"
)

type IUserService interface {
	GetById(id int) (models.MinimalUserModel, error)
	GetAll() []models.MinimalUserModel
	Delete(id int) bool
}

type UserService struct {
	Repo repositories.Repository
}

func NewUserService(repo repositories.Repository) IUserService {
	return &UserService{repo}
}

func (s *UserService) GetById(id int) (models.MinimalUserModel, error) {
	user, err := s.Repo.UserRepostory.GetById(id)
	minUser := &models.MinimalUserModel{
		Id: user.Id,
		Fullname: user.Fullname,
		Email: user.Email,
		RoleId: user.RoleId,
	}

	return *minUser, err
}

func (s *UserService) GetAll() []models.MinimalUserModel {
	users, _ := s.Repo.UserRepostory.GetAll()
	var minUsers []models.MinimalUserModel
	var minUser models.MinimalUserModel
	
	for _, user := range users {
		minUser = models.MinimalUserModel{
			Id: user.Id,
			Fullname: user.Fullname,
			Email: user.Email,
			RoleId: user.RoleId,
		}
		minUsers = append(minUsers, minUser)
	}

	return minUsers
}

func (s *UserService) Delete(id int) bool {
	s.Repo.UserRepostory.Delete(id)
	return true
}


