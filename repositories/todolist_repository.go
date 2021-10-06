package repositories

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
)

type ITodolistRepository interface {
	Create(todolist models.CreateTodolistModel) error
	Update(id int, todolist models.UpdateTodolistModel) error
	GetAll() ([]base.TodoList, error)
	GetById(id int) (base.TodoList, error)
	Delete(id int) error
}

type TodolistRepository struct {
	db *pg.DB
}

func NewTodolistRepository(db *pg.DB) ITodolistRepository {
	return &TodolistRepository{db}
}

func (r *TodolistRepository) Create(todolist models.CreateTodolistModel) (error) {
	newTodolist := &base.TodoList{
		IdTodo: todolist.IdTodo,
		Task: todolist.Task,
		IdStatus: todolist.IdStatus,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := r.db.Model(newTodolist).Insert()
	return err
}

func (r *TodolistRepository) GetAll() ([]base.TodoList, error) {
	var todolists []base.TodoList
	err := r.db.Model(&todolists).Select()
	return todolists, err
}

func (r *TodolistRepository) GetById(id int) (base.TodoList, error) {
	var todolist base.TodoList
	err := r.db.Model(&todolist).Where("id = ?", id).Select()
	return todolist, err
}

func (r *TodolistRepository) Update(id int, todolist models.UpdateTodolistModel) error {
	values := map[string]interface{}{
		"task": todolist.Task,
		"id_status": todolist.IdStatus,
	}

	res, err := r.db.Model(&values).TableExpr("todolist").Where("id = ?", id).Update()
	log.Print(res)
	return err
}

func (r *TodolistRepository) Delete(id int) error {
	var todolist base.TodoList
	_, err := r.db.Model(&todolist).Where("id = ?", id).Delete()
	return err
}

