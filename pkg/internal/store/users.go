package store

import (
	"context"
	"database/sql"

	"github.com/Xanssun/technical-specification.git/pkg/internal/models"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (username, password, email)
		VALUES($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
