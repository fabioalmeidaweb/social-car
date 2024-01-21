package models

import (
	"errors"
	"social-car/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user used in the social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare prepares a user
func (user *User) Prepare(step string) error {
	if err := user.Validate(step); err != nil {
		return err
	}

	if err := user.Format("register"); err != nil {
		return err
	}
	return nil
}

// Validate validates a user data
func (user *User) Validate(step string) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}

	if user.Nick == "" {
		return errors.New("Nick is required")
	}

	if user.Email == "" {
		return errors.New("Email is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Invalid email format")
	}

	if step == "register" && user.Password == "" {
		return errors.New("Password is required")
	}

	return nil
}

// Format formats a user
func (user *User) Format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		hashPass, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPass)
	}

	return nil
}
