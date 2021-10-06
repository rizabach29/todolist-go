package services

import (
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
	"github.com/rizabach29/todolist-go/repositories"
)

type ITodoService interface {
	Create(newTodo models.CreateTodoModel) error
	Update(id int, updatedTodo models.UpdateTodoModel) error
	GetById(id int) (base.Todo, error)
	GetAll() ([]base.Todo, error)
	Delete(id int) error
}

type TodoService struct {
	Repo repositories.Repository
}

func NewTodoService(Repo repositories.Repository) ITodoService {
	return &TodoService{Repo}
}

func (s *TodoService) Create(m models.CreateTodoModel) error{
	err := s.Repo.TodoRepository.Create(m)
	return err
}

func (s *TodoService) Update(id int, m models.UpdateTodoModel) error{
	err := s.Repo.TodoRepository.Update(id, m)
	return err
}

func (s *TodoService) Delete(id int) error{
	err := s.Repo.TodoRepository.Delete(id)
	return err
}

func (s *TodoService) GetById(id int) (base.Todo, error){
	todo, err := s.Repo.TodoRepository.GetById(id)
	return todo, err
}

func (s *TodoService) GetAll() ([]base.Todo, error){
	todos, err := s.Repo.TodoRepository.GetAll()
	return todos, err
}