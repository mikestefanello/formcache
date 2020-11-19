package router

import (
	"github.com/go-chi/chi"
	"github.com/mikestefanello/formcache/handlers"
)

func NewRouter(h *handlers.HTTPHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/steps", h.Steps)
	r.Post("/steps", h.Steps)
	return r
}
