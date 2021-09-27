package services

import (
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/repositories"
)

type ITodoService interface {
	Create(newTodo models.CreateTodoModel) models.Todo
	Update(updatedTodo models.UpdateTodoModel) models.Todo
	GetById(id int) models.Todo
	GetAll() []models.Todo
	Delete() bool
}

type TodoService struct {
}

func NewTodoService(repo repositories.Repository) ITodoService {
	return &TodoService{}
}

func (s *TodoService) Create(m models.CreateTodoModel) models.Todo{
	return models.Todo{}
}

func (s *TodoService) Update(m models.UpdateTodoModel) models.Todo{
	return models.Todo{}
}

func (s *TodoService) Delete() bool{
	return true
}

func (s *TodoService) GetById(id int) models.Todo{
	return models.Todo{}
}

func (s *TodoService) GetAll() []models.Todo{
	return []models.Todo{}
}