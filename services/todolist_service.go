package services

import (
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
	"github.com/rizabach29/todolist-go/repositories"
)

type ITodolistService interface {
	Create(todolist models.CreateTodolistModel) error
	GetAll() ([]base.TodoList, error)
	GetById(id int) (base.TodoList, error)
	GetByTodoId(id int) (base.TodoList, error)
	Update(id int, todolist models.UpdateTodolistModel) error
	Delete(id int) error
}

type TodolistService struct {
	repo repositories.Repository
}

func NewTodolistService(repo repositories.Repository) ITodolistService {
	return &TodolistService{repo}
}

func (s *TodolistService) Create(todolist models.CreateTodolistModel) error {
	err := s.repo.TodolistRepository.Create(todolist)
	return err
}

func (s *TodolistService) GetAll() ([]base.TodoList, error) {
	todolists, err := s.repo.TodolistRepository.GetAll()
	return todolists, err
}

func (s *TodolistService) GetById(id int) (base.TodoList, error) {
	todolist, err := s.repo.TodolistRepository.GetById(id)
	return todolist, err
}

func (s *TodolistService) GetByTodoId(id int) (base.TodoList, error) {
	todolist, err := s.repo.TodolistRepository.GetByTodoId(id)
	return todolist, err
}

func (s *TodolistService) Update(id int, todolist models.UpdateTodolistModel) error {
	err := s.repo.TodolistRepository.Update(id, todolist)
	return err
}

func (s *TodolistService) Delete(id int) error {
	err := s.repo.TodolistRepository.Delete(id)
	return err
}

