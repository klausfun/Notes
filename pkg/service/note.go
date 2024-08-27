package service

import (
	"Notes/models"
	"Notes/pkg/repository"
)

type NoteService struct {
	repo repository.Notes
}

func NewNoteService(repo repository.Notes) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(userId int, note models.Note) (int, error) {
	return s.repo.Create(userId, note)
}

func (s *NoteService) GetAll(userId int) ([]models.Note, error) {
	return s.repo.GetAll(userId)
}
