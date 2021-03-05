package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// User struct describe user object.
type User struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Email      string    `db:"email" json:"email" validate:"required,email"`
	UserStatus int       `db:"user_status" json:"user_status"`
	UserAttrs  UserAttrs `db:"user_attrs" json:"user_attrs"`
}

// UserAttrs struct describe user attributes.
type UserAttrs struct {
	Picture   string `json:"picture"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	About     string `json:"about"`
}

// Value make the UserAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (u UserAttrs) Value() (driver.Value, error) {
	return json.Marshal(u)
}

// Scan make the UserAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (u *UserAttrs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &u)
}
