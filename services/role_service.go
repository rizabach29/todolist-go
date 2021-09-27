package services

import (
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/repositories"
)

type IRoleService interface {
	Create(newRole models.CreateRoleModel) models.Role
	Update(updatedRole models.UpdateRoleModel) models.Role
	GetById(id int) models.Role
	GetAll() []models.Role
	Delete() bool
}

type RoleService struct {
}

func NewRoleService(repo repositories.Repository) IRoleService {
	return &RoleService{}
}

func (s *RoleService) Create(m models.CreateRoleModel) models.Role{
	return models.Role{}
}

func (s *RoleService) Update(m models.UpdateRoleModel) models.Role{
	return models.Role{}
}

func (s *RoleService) Delete() bool{
	return true
}

func (s *RoleService) GetById(id int) models.Role{
	return models.Role{}
}

func (s *RoleService) GetAll() []models.Role{
	return []models.Role{}
}