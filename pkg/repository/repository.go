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
	Create(userId int, note models.Note) (int, error)
	GetAll(userId int) ([]models.Note, error)
}

type Repository struct {
	Authorization
	Notes
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Notes:         NewNotePostgres(db),
	}
}
