package repositories

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/app"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
)

type ITodoRepository interface {
	Create(todo models.CreateTodoModel) error
	Update(id int, todo models.UpdateTodoModel) error
	GetAll() ([]base.Todo, error)
	GetById(id int) (base.Todo, error)
	Delete(id int) error
}

type TodoRepository struct{
	db *pg.DB
}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{app.GetDatabase()}
}

func (repo *TodoRepository) Create(todo models.CreateTodoModel) error{
	newTodo := &base.Todo{
		Title: todo.Title,
		UserId: todo.UserId,
		CreatedAt : time.Now(),
		UpdatedAt : time.Now(),
	}

	_, err := repo.db.Model(newTodo).Insert()
	return err
}

func (repo *TodoRepository) Update(id int, todo models.UpdateTodoModel) error {
	values := map[string]interface{}{
		"title": todo.Title,
	}
	
	_, err := repo.db.Model(&values).TableExpr("todos").Where("id = ?", id).Update()
	return err
}

func (repo *TodoRepository) Delete(id int) error {
	model := new(base.Todo)
	_, err := repo.db.Model(model).Where("id = ?", id).Delete()
	return err
}

func (repo *TodoRepository) GetById(id int) (base.Todo, error) {
	var todo base.Todo
	err := repo.db.Model(&todo).Where("id = ?", id).Select()
	log.Print(todo)
	return todo, err
}

func (repo *TodoRepository) GetAll() ([]base.Todo, error) {
	var todos []base.Todo

	err := repo.db.Model(&todos).Select()
	return todos, err
}