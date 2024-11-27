package store

import (
	"context"
	"database/sql"

	"github.com/Xanssun/technical-specification.git/pkg/internal/models"
)

type Storage struct {
	Users interface {
		Create(context.Context, *models.User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UsersStore{db},
	}
}
