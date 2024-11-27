package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Xanssun/technical-specification.git/pkg/internal/models"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Users interface {
		Create(context.Context, *models.User) error
		GetByID(context.Context, int64) (*models.User, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Users: &UsersStore{db},
	}
}
