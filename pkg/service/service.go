package service

import (
	"Notes/models"
	"Notes/pkg/repository"
)

const (
	spellCheckUrl = "https://speller.yandex.net/services/spellservice.json/checkTexts"
)

type Authorization interface {
	CreateUser(user models.Users) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Notes interface {
	Create(userId int, note models.Note) (int, error)
	GetAll(userId int) ([]models.Note, error)
}

type SpellCheck interface {
	SpellChecking(spellCheck [][]SpellCheckResult, text string) (string, error)
	GettingErrors(text string) ([][]SpellCheckResult, error)
}

type Service struct {
	Authorization
	Notes
	SpellCheck
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Notes:         NewNoteService(repos.Notes),
		SpellCheck:    NewSpellCheckService(spellCheckUrl),
	}
}
