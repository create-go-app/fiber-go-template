package queries

import (
	"github.com/create-go-app/fiber-go-template/app/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
	*sqlx.DB
}

// FindUserByID ...
func (s *UserQueries) FindUserByID(id uuid.UUID) (models.User, error) {
	var user models.User
	if err := s.Get(&user, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUsers ...
func (s *UserQueries) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := s.Select(&users, `SELECT * FROM users`); err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// CreateUser ...
func (s *UserQueries) CreateUser(u *models.User) error {
	if _, err := s.Exec(
		`INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6)`,
		u.ID,
		u.CreatedAt,
		u.UpdatedAt,
		u.Email,
		u.UserStatus,
		u.UserAttrs,
	); err != nil {
		return err
	}
	return nil
}

// UpdateUser ...
func (s *UserQueries) UpdateUser(u *models.User) error {
	if _, err := s.Exec(
		`UPDATE users SET updated_at = $2, email = $3, user_attrs = $4 WHERE id = $1`,
		u.ID,
		u.UpdatedAt,
		u.Email,
		u.UserAttrs,
	); err != nil {
		return err
	}
	return nil
}

// DeleteUser ...
func (s *UserQueries) DeleteUser(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM users WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
