package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// user on social media
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"senha,omitempty"`
	CreatedIn time.Time `json:"CreatedIn,omitempty"`
}

// calls validate and format
func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	if error := user.format(step); error != nil {
		return error
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}
	if user.Nick == "" {
		return errors.New("Nick is required")
	}
	if user.Email == "" {
		return errors.New("Email is required")
	}
	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("Inserted email invalid")
	}
	if step == "registration" && user.Password == "" {
		return errors.New("A password is required")
	}
	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "registration" {
		passWHash, error := security.Hash(user.Password)
		if error != nil {
			return error
		}
		user.Password = string(passWHash)
	}
	return nil
}
