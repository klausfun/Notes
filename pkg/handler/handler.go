package handler

import (
	"Notes/pkg/service"
	"github.com/go-chi/chi"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/auth", func(r chi.Router) {
		r.Post("/signUp", h.signUp)
		r.Post("/signIn", h.signIn)
	})

	router.Route("/api", func(r chi.Router) {
		r.Use(h.userIdentity)

		r.Route("/notes", func(r chi.Router) {
			r.Post("/", h.createNote)
			r.Get("/", h.getAllNotes)
		})

	})

	return router
}
