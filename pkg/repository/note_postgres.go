package repository

import (
	"Notes/models"
	"database/sql"
	"fmt"
)

type NotePostgres struct {
	db *sql.DB
}

func NewNotePostgres(db *sql.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (r *NotePostgres) Create(userId int, note models.Note) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, title, description)"+
		" VALUES ($1, $2, $3) RETURNING id", notesTable)
	row := r.db.QueryRow(query, userId, note.Title, note.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *NotePostgres) GetAll(userId int) ([]models.Note, error) {
	var notes []models.Note
	query := fmt.Sprintf("SELECT nt.id, nt.title, nt.description FROM %s nt INNER JOIN %s us on us.id = nt.user_id"+
		" WHERE us.id = $1", notesTable, userTable)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var note models.Note
		if err = rows.Scan(&note.Id, &note.Title, &note.Description); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, rows.Err()
}
