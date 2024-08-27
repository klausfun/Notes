package handler

import (
	"Notes/models"
	"encoding/json"
	"net/http"
)

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	if err != nil {
		return
	}

	var input models.Note
	if err = json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, "invalid input body")
		return
	}

	spellChecks, err := h.services.GettingErrors(input.Description)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, "spelling error")
		return
	}

	formattedText, err := h.services.SpellChecking(spellChecks, input.Description)
	if err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	input.Description = formattedText
	id, err := h.services.Create(userId, input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

type getAllNotesResponse struct {
	Data []models.Note `json:"data"`
}

func (h *Handler) getAllNotes(w http.ResponseWriter, r *http.Request) {
	userId, err := GetUserId(r)
	if err != nil {
		return
	}

	notes, err := h.services.GetAll(userId)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(getAllNotesResponse{
		Data: notes,
	})
}
