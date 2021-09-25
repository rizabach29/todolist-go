package services

import "github.com/rizabach29/todolist-go/repositories"


type Services struct {
	UserService IAuthService
}

// Initialize all services
func NewService(repo *repositories.Repository) *Services {
	return &Services{
		UserService: NewAuthService(*repo),
	}
}