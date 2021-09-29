package repositories

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/models"
)

type IUserRepository interface {
	Create(newUser models.RegisterModel) models.RegisterModel
	Update() models.User
	Delete(id int)
	GetById(email string) (models.User, error)
	GetAll() []models.User
}

type UserRepository struct{
	db *pg.DB
}

func NewUserRepository(db *pg.DB) IUserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) Create(user models.RegisterModel) models.RegisterModel {
	
	newUser := &models.User{
		Fullname: user.Fullname,
		Email: user.Email,
		Password: user.Password,
		ForgotPassword: user.ForgotPassword,
		RoleId: user.RoleId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_,err := repo.db.Model(newUser).Insert()
	helpers.PanicIfError(err)

	return user
}

func (repo *UserRepository) Update() models.User {
	return models.User{}
}

func (repo *UserRepository) Delete(id int) {
	model := new(models.User)
	_, err := repo.db.Model(model).Where("id = ?", id).Delete()
	helpers.PanicIfError(err)
}

func (repo *UserRepository) GetById(email string) (models.User, error) {
	var user models.User
	err := repo.db.Model(&user).Where("email = ?", email).Select()
	log.Print(user)
	return user, err
}

func (repo *UserRepository) GetAll() []models.User {
	var users []models.User

	err := repo.db.Model(&users).Select()
	helpers.PanicIfError(err)

	return users
}