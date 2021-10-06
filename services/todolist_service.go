package services

import (
	"github.com/rizabach29/todolist-go/models"
	"github.com/rizabach29/todolist-go/models/base"
	"github.com/rizabach29/todolist-go/repositories"
)

type ITodolistService interface {
	Create(todolist models.CreateTodolistModel) base.TodoList
	GetAll() []base.TodoList
	GetById(id int) (base.TodoList, error)
	Update(todolist models.UpdateTodolistModel) base.TodoList
	Delete(id int)
}

type TodolistService struct {
	repo *repositories.Repository
}

func NewTodolistService(repo *repositories.Repository) ITodolistService {
	return &TodolistService{repo}
}

func (s *TodolistService) Create(todolist models.CreateTodolistModel) base.TodoList {
	panic("not implemented") // TODO: Implement
}

func (s *TodolistService) GetAll() []base.TodoList {
	panic("not implemented") // TODO: Implement
}

func (s *TodolistService) GetById(id int) (base.TodoList, error) {
	panic("not implemented") // TODO: Implement
}

func (s *TodolistService) Update(todolist models.UpdateTodolistModel) base.TodoList {
	panic("not implemented") // TODO: Implement
}

func (s *TodolistService) Delete(id int) {
	panic("not implemented") // TODO: Implement
}

