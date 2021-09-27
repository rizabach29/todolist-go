package repositories

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/models"
)

type IRoleRepository interface {
	Create(Role models.CreateRoleModel) models.Role
	Update(Role models.CreateRoleModel) models.Role
	Delete(id int)
	GetById(id int) (models.Role, error)
	GetAll() []models.Role
}

type RoleRepository struct{
	db *pg.DB
}

func NewRoleRepository(db *pg.DB) IRoleRepository {
	return &RoleRepository{db}
}

func (repo *RoleRepository) Create(role models.CreateRoleModel) models.Role {
	newRole := &models.Role{
		Name: role.Name,
		Description: role.Description,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}
	
	_, err := repo.db.Model(newRole).Insert()
	helpers.PanicIfError(err)
	
	return *newRole
}

func (repo *RoleRepository) Update(role models.CreateRoleModel) models.Role {
	newRole := &models.Role{
		Name: role.Name,
		Description: role.Description,
		UpdatedAt : time.Now(),
	}
	
	_, err := repo.db.Model(newRole).WherePK().Update()
	helpers.PanicIfError(err)
	
	return *newRole
}

func (repo *RoleRepository) Delete(id int) {
	model := new(models.Role)
	_, err := repo.db.Model(model).Where("id = ?", id).Delete()
	helpers.PanicIfError(err)
}

func (repo *RoleRepository) GetById(id int) (models.Role, error) {
	var role models.Role
	err := repo.db.Model(role).Where("id = ?", id).Select()
	helpers.PanicIfError(err)

	return role, err
}

func (repo *RoleRepository) GetAll() []models.Role {
	var roles []models.Role

	err := repo.db.Model(&roles).Select()
	helpers.PanicIfError(err)

	return roles
}