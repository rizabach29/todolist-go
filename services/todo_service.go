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
	Delete(id int) bool
}

type TodoService struct {
	Repo repositories.Repository
}

func NewTodoService(Repo repositories.Repository) ITodoService {
	return &TodoService{Repo}
}

func (s *TodoService) Create(m models.CreateTodoModel) models.Todo{
	newTodo := s.Repo.TodoRepository.Create(m)
	return newTodo
}

func (s *TodoService) Update(m models.UpdateTodoModel) models.Todo{
	updatedTodo := s.Repo.TodoRepository.Update(m)
	return updatedTodo
}

func (s *TodoService) Delete(id int) bool{
	s.Repo.TodoRepository.Delete(id)
	return true
}

func (s *TodoService) GetById(id int) models.Todo{
	todo, _ := s.Repo.TodoRepository.GetById(id)
	return todo
}

func (s *TodoService) GetAll() []models.Todo{
	todos := s.Repo.TodoRepository.GetAll()
	return todos
}