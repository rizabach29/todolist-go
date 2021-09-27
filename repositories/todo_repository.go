package repositories

import (
	"time"

	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/helpers"
	"github.com/rizabach29/todolist-go/models"
)

type ITodoRepository interface {
	Create(todo models.CreateTodoModel) models.Todo
	Update(todo models.CreateTodoModel) models.Todo
	Delete(id int)
	GetById(id int) (models.Todo, error)
	GetAll() []models.Todo
}

type TodoRepository struct{
	db *pg.DB
}

func NewTodoRepository(db *pg.DB) ITodoRepository {
	return &TodoRepository{db}
}

func (repo *TodoRepository) Create(todo models.CreateTodoModel) models.Todo {
	newTodo := &models.Todo{
		Title: todo.Title,
		UserId: todo.UserId,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}

	_, err := repo.db.Model(newTodo).Insert()
	helpers.PanicIfError(err)

	return *newTodo
}

func (repo *TodoRepository) Update(todo models.CreateTodoModel) models.Todo {
	newTodo := &models.Todo{
		Title: todo.Title,
		UserId: todo.UserId,
		UpdatedAt : time.Now(),
	}
	
	_, err := repo.db.Model(newTodo).WherePK().Update()
	helpers.PanicIfError(err)
	
	return *newTodo
}

func (repo *TodoRepository) Delete(id int) {
	model := new(models.Todo)
	_, err := repo.db.Model(model).Where("id = ?", id).Delete()
	helpers.PanicIfError(err)
}

func (repo *TodoRepository) GetById(id int) (models.Todo, error) {
	var todo models.Todo
	err := repo.db.Model(todo).Where("id = ?", id).Select()
	helpers.PanicIfError(err)

	return todo, err
}

func (repo *TodoRepository) GetAll() []models.Todo {
	var todos []models.Todo

	err := repo.db.Model(&todos).Select()
	helpers.PanicIfError(err)

	return todos
}