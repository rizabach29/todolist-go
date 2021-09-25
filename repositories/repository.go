package repositories

import (
	"github.com/go-pg/pg"
)

type Repository struct {
	UserRepostory IUserRepository
}

// initialize all repositories
func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		UserRepostory: NewUserRepository(db),
	}
}