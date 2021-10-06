package services

import (
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
	"github.com/rizabach29/todolist-go/repositories"
)

type IRoleService interface {
	Create(newRole models.CreateRoleModel) base.Role
	Update(updatedRole models.UpdateRoleModel) base.Role
	GetById(id int) base.Role
	GetAll() []base.Role
	Delete() bool
}

type RoleService struct {
}

func NewRoleService(repo repositories.Repository) IRoleService {
	return &RoleService{}
}

func (s *RoleService) Create(m models.CreateRoleModel) base.Role{
	return base.Role{}
}

func (s *RoleService) Update(m models.UpdateRoleModel) base.Role{
	return base.Role{}
}

func (s *RoleService) Delete() bool{
	return true
}

func (s *RoleService) GetById(id int) base.Role{
	return base.Role{}
}

func (s *RoleService) GetAll() []base.Role{
	return []base.Role{}
}