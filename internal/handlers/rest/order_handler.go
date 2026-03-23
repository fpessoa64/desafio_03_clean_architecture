package rest

import (
	"encoding/json"
	"net/http"

	"github.com/fpessoa64/desafio03_clean_arch/internal/entity"
	"github.com/fpessoa64/desafio03_clean_arch/internal/usecase"
)

type Handler struct {
	UC *usecase.OrderUsecase
}

func NewHandler(uc *usecase.OrderUsecase) *Handler {
	return &Handler{UC: uc}
}

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var o entity.Order
	if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.UC.Create(r.Context(), &o); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.UC.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}
