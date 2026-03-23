package routes

import (
	"github.com/fpessoa64/desafio03_clean_arch/internal/handlers/rest"
	"github.com/go-chi/chi/v5"
)

func RegisterOrderRoutes(r chi.Router, h *rest.Handler) {
	r.Route("/order", func(r chi.Router) {
		r.Post("/", h.CreateOrder)
		r.Get("/", h.ListOrders)
	})
}
