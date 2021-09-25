package repositories

import (
	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/models"
)

type IUserRepository interface {
	Create(newUser models.RegisterModel) models.User
	Update() models.User
	Delete()
	GetById(email string) (models.User, error)
	GetAll() []models.User
}

type UserRepository struct{
	db *pg.DB
}

func NewUserRepository(db *pg.DB) IUserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(newUser models.RegisterModel) models.User {
	
	return models.User{}
}

func (repo *UserRepository) Update() models.User {
	return models.User{}
}

func (repo *UserRepository) Delete() {
	
}

func (repo *UserRepository) GetById(email string) (models.User, error) {
	return models.User{}, nil
}

func (repo *UserRepository) GetAll() []models.User {
	return []models.User{}
}