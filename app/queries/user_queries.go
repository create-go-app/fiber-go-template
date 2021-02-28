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

// GetUsers func for getting all users.
func (q *UserQueries) GetUsers() ([]models.User, error) {
	// Define users variable.
	var users []models.User

	// Send query to database.
	if err := q.Select(&users, `SELECT * FROM users`); err != nil {
		return []models.User{}, err
	}

	return users, nil
}

// GetUser func for getting one user by given ID.
func (q *UserQueries) GetUser(id uuid.UUID) (models.User, error) {
	// Define user variable.
	var user models.User

	// Send query to database.
	if err := q.Get(&user, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

// CreateUser func for creating user by given User object.
func (q *UserQueries) CreateUser(u *models.User) error {
	// Send query to database.
	if _, err := q.Exec(
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
func (q *UserQueries) UpdateUser(u *models.User) error {
	// Send query to database.
	if _, err := q.Exec(
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
func (q *UserQueries) DeleteUser(id uuid.UUID) error {
	// Send query to database.
	if _, err := q.Exec(`DELETE FROM users WHERE id = $1`, id); err != nil {
		return err
	}

	return nil
}
