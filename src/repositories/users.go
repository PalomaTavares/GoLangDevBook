package repositories

import (
	"api/src/models"
	"database/sql"
)

// repositoru of users
type users struct {
	db *sql.DB
}

// creates user repository
func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uint64, error) {
	return 0, nil
}
