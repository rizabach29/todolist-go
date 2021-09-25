package repositories

import "github.com/rizabach29/todolist-go/models"

type ITodoRepository interface {
	Create() models.Todo
	Update() models.Todo
	Delete()
	GetById() models.Todo
	GetAll() []models.Todo
}

type TodoRepository struct{}

func NewTodoRepository() ITodoRepository {
	return &TodoRepository{}
}

func (repo *TodoRepository) Create() models.Todo {
	return models.Todo{}
}

func (repo *TodoRepository) Update() models.Todo {
	return models.Todo{}
}

func (repo *TodoRepository) Delete() {
	
}

func (repo *TodoRepository) GetById() models.Todo {
	return models.Todo{}
}

func (repo *TodoRepository) GetAll() []models.Todo {
	return []models.Todo{}
}