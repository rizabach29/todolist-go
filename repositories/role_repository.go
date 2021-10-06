package repositories

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
)

type IRoleRepository interface {
	Create(Role models.CreateRoleModel) error
	Update(id int, Role models.CreateRoleModel) error
	Delete(id int) error
	GetById(id int) (base.Role, error)
	GetAll() ([]base.Role, error)
}

type RoleRepository struct{
	db *pg.DB
}

func NewRoleRepository(db *pg.DB) IRoleRepository {
	return &RoleRepository{db}
}

func (repo *RoleRepository) Create(role models.CreateRoleModel) error {
	newRole := &base.Role{
		Name: role.Name,
		Description: role.Description,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	
	_, err := repo.db.Model(newRole).Insert()
	return err
}

func (repo *RoleRepository) Update(id int, role models.CreateRoleModel) error {
	values := map[string]interface{}{
		"name": role.Name,
		"description": role.Description,
	}
	
	_, err := repo.db.Model(&values).TableExpr("roles").Where("id = ?", id).Update()
	return err
}

func (repo *RoleRepository) Delete(id int) error {
	model := new(base.Role)
	_, err := repo.db.Model(model).Where("id = ?", id).Delete()
	return err
}

func (repo *RoleRepository) GetById(id int) (base.Role, error) {
	var role base.Role
	err := repo.db.Model(&role).Where("id = ?", id).Select()

	return role, err
}

func (repo *RoleRepository) GetAll() ([]base.Role, error) {
	var roles []base.Role

	err := repo.db.Model(&roles).Select()
	return roles, err
}