package models

type Users struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Password string `json:"password" binding:"required" db:"password_hash"`
	Email    string `json:"email" binding:"required" db:"email"`
}
