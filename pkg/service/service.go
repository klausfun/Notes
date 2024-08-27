package service

import (
	"Notes/models"
	"Notes/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.Users) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Notes interface {
}

type Service struct {
	Authorization
	Notes
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
