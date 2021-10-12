package repositories

type Repository struct {
	UserRepostory IUserRepository
	RoleRepository IRoleRepository
	TodoRepository ITodoRepository
	StatusRepository IStatusRepository
	AttachmentRepository IAttachmentRepository
	TodolistRepository ITodolistRepository
}

// initialize all repositories
func NewRepository() *Repository {
	return &Repository{
		UserRepostory: NewUserRepository(),
		RoleRepository: NewRoleRepository(),
		TodoRepository: NewTodoRepository(),
		StatusRepository: NewStatusRepository(),
		AttachmentRepository: NewAttachmentRepository(),
		TodolistRepository: NewTodolistRepository(),
	}
}