package repository

import (
	"Notes/models"
	"database/sql"
	"fmt"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.Users) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, password_hash, email)"+
		" VALUES ($1, $2, $3) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Name, user.Password, user.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.Users, error) {
	var user models.Users
	query := fmt.Sprintf("SELECT id FROM %s "+
		"WHERE email = $1 AND password_hash = $2", userTable)
	err := r.db.QueryRow(query, email, password).Scan(&user.Id)

	return user, err
}
