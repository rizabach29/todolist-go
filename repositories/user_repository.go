package repositories

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/app"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
)

type IUserRepository interface {
	Create(newUser models.RegisterModel) error
	Update(id int, user models.UpdateUserModel) error
	Delete(id int) error
	GetById(id int) (base.User, error)
	GetByEmail(email string) (base.User, error)
	GetAll() ([]base.User, error)
}

type UserRepository struct{
	db *pg.DB
}

func NewUserRepository() IUserRepository {
	return &UserRepository{app.GetDatabase()}
}

func (repo *UserRepository) Create(user models.RegisterModel) error {
	
	newUser := &base.User{
		Fullname: user.Fullname,
		Email: user.Email,
		Password: user.Password,
		ForgotPassword: user.ForgotPassword,
		RoleId: user.RoleId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := repo.db.Model(newUser).Insert()
	return err
}

func (repo *UserRepository) Update(id int, user models.UpdateUserModel) error {
	values := &base.User{
		Id: id,
		Fullname: user.Fullname,
		RoleId: user.RoleId,
	}
	log.Print(values)
	
	_, err := repo.db.Model(values).
				Set("fullname = ?fullname").
				Set("role_id = ?role_id").WherePK().Update()
	log.Print(err)
	return err
}

func (repo *UserRepository) Delete(id int) error {
	model := new(base.User)
	_, err := repo.db.Model(model).Where("id = ?", id).Delete()
	return err
}

func (repo *UserRepository) GetByEmail(email string) (base.User, error) {
	var user base.User
	err := repo.db.Model(&user).Where("email = ?", email).Select()
	return user, err
}

func (repo *UserRepository) GetById(id int) (base.User, error) {
	var user base.User
	err := repo.db.Model(&user).Where("id = ?", id).Select()
	return user, err
}

func (repo *UserRepository) GetAll() ([]base.User, error) {
	var users []base.User
	err := repo.db.Model(&users).Select()
	return users, err
}