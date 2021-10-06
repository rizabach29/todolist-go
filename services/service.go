package services

import "github.com/rizabach29/todolist-go/repositories"


type Services struct {
	AuthService IAuthService
	TodoService ITodoService
	RoleService IRoleService
	UserService IUserService
}

// Initialize all services
func NewService(repo *repositories.Repository) *Services {
	return &Services{
		AuthService: NewAuthService(*repo),
		TodoService: NewTodoService(*repo),
		RoleService: NewRoleService(*repo),
		UserService: NewUserService(*repo),
	}
}