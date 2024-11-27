package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Xanssun/technical-specification.git/pkg/internal/models"
)

var (
	ErrDuplicateEmail    = errors.New("a user with that email already exists")
	ErrDuplicateUsername = errors.New("a user with that username already exists")
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *models.User) error {
	// Хэшируем пароль
	hashedPassword, err := models.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	query := `
		INSERT INTO users (username, password, email)
		VALUES($1, $2, $3) RETURNING id, created_at
	`

	err = s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password, // Хэшированный пароль
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (s *UsersStore) GetByID(ctx context.Context, userID int64) (*models.User, error) {
	query := `
		SELECT id, username, email, password, created_at
		FROM users
		WHERE id = $1
	`

	user := &models.User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password, // Получаем хэш пароля
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return user, nil
}
