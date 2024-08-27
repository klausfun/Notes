package repository

import (
	"Notes/models"
	"database/sql"
)

type Authorization interface {
	CreateUser(user models.Users) (int, error)
	GetUser(email, password string) (models.Users, error)
}

type Notes interface {
}

type Repository struct {
	Authorization
	Notes
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
