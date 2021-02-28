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

// FindUserByID func for searching user by given ID.
func (s *UserQueries) FindUserByID(id uuid.UUID) (models.User, error) {
	// Define user variable.
	var user models.User

	// Send query to database.
	if err := s.Get(&user, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetUsers func for getting all exists users.
func (s *UserQueries) GetUsers() ([]models.User, error) {
	// Define users variable.
	var users []models.User

	// Send query to database.
	if err := s.Select(&users, `SELECT * FROM users`); err != nil {
		return []models.User{}, err
	}

	return users, nil
}

// CreateUser func for creating user by given User object.
func (s *UserQueries) CreateUser(u *models.User) error {
	// Send query to database.
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

// UpdateUser func for updating user by given User object.
func (s *UserQueries) UpdateUser(u *models.User) error {
	// Send query to database.
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

// DeleteUser func for delete user by given ID.
func (s *UserQueries) DeleteUser(id uuid.UUID) error {
	// Send query to database.
	if _, err := s.Exec(`DELETE FROM users WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}
